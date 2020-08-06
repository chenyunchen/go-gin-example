package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

var tplMap = sync.Map{}

// `,"status":"{{ index .query.status 0 }}"` +
// `,"flashid":"{{.body.flashid}}"` +
var defaultTpl = Template{
	Resp: `{{.reqBody}}`,
	CallbackURL: `{{ index .query.notify_url 0 }}` +
		`?time={{ index .query.time 0 }}` +
		`&sign={{ index .query.sign 0 }}`,
	CallbackBody: `{` +
		`"flashid":"{{.body.flashid}}"` +
		`,"status":"{{ index .query.status 0 }}"` +
		`}`,
}

type Template struct {
	Resp         string
	CallbackURL  string
	CallbackBody string
}

func main() {
	e := echo.New()

	e.POST("/add/:target", func(c echo.Context) error {
		t := Template{}
		_ = c.Bind(&t)
		fmt.Println("add", t)
		tplMap.Store(c.Param("target"), t)
		return c.NoContent(http.StatusNoContent)
	})

	e.Any("/mock/:target", func(c echo.Context) error {
		tplInterface, ok := tplMap.Load(c.Param("target"))
		if !ok {
			tplInterface = defaultTpl
		}
		tpl := tplInterface.(Template)
		respTpl, _ := template.New("resp").Parse(tpl.Resp)
		callbackURLTpl, _ := template.New("callbackURL").Parse(tpl.CallbackURL)
		callbackBodyTpl, _ := template.New("callbackBody").Parse(tpl.CallbackBody)

		b, _ := ioutil.ReadAll(c.Request().Body)
		body := map[string]interface{}{}
		_ = json.Unmarshal(b, &body)

		query := c.Request().URL.Query()
		delay, _ := strconv.Atoi(query.Get("delay"))
		retry, _ := strconv.Atoi(query.Get("retry"))

		m := map[string]interface{}{
			"query":   query,
			"body":    body,
			"reqBody": string(b),
		}
		resp := bytes.Buffer{}
		_ = respTpl.Execute(&resp, m)

		// callback
		go func() {
			for ; retry >= 0; retry-- {
				callbackURL := bytes.Buffer{}
				_ = callbackURLTpl.Execute(&callbackURL, m)

				callbackBody := bytes.Buffer{}
				_ = callbackBodyTpl.Execute(&callbackBody, m)

				time.Sleep(time.Duration(delay) * time.Second)
				u, _ := url.Parse(callbackURL.String())
				fmt.Println("callback", u)
				resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(callbackBody.Bytes()))
				fmt.Println("resp", resp, err)
			}
		}()

		// response
		status, _ := strconv.Atoi(query.Get("status"))
		if status == 0 {
			status = http.StatusOK
		}
		w := c.Response().Writer
		w.WriteHeader(status)
		_, _ = w.Write(resp.Bytes())
		return nil
	})
	e.Logger.Fatal(e.Start(":8080"))
}
