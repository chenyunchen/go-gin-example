package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create", handler)
	r.POST("/callback", callbackhandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handler(c *gin.Context) {
	var data PolypayDepositeRequest
	timeStramp := c.Query("time")
	//sign := c.Query("sign")

	err := c.BindJSON(&data)
	if err != nil {
		log.Println("BindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid",
		})
		return
	}
	uid := xid.New().String()
	resp := SuccessResp{
		Status:        1000,
		Message:       "success",
		TradeID:       data.MerchantID,
		OrderID:       data.OrderID,
		PayURL:        "this.PayURL",
		PayeeRealName: "Heyhey",
	}

	d, _ := json.Marshal(resp)
	c.Data(http.StatusAccepted, "", d)
	go func() {
		time.Sleep(2 * time.Second)
		call := CallBackReq{
			Status:        1000,
			TradeAmount:   data.BillPrice,
			ReceiptAmount: data.BillPrice,
			TradeID:       data.OrderID,
			UpstreamOrder: uid,
			RepeatPay:     false,
			OrderID:       data.MerchantID,
		}
		callBack, err := json.Marshal(call)
		if err != nil {
			log.Println("Marshal", err)
			return
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", data.NotifyURL, strings.NewReader(string(callBack)))
		if err != nil {
			log.Println("NewRequest", err)
		}
		apiKey := "brsm0oofvt9ck0mpd520"
		m := md5.New()
		m.Write([]byte(timeStramp + apiKey))
		s := hex.EncodeToString(m.Sum(nil))
		req.Header.Add("X-Signature", s)
		req.Header.Add("X-Timestamp", timeStramp)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("client.Do", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil.ReadAll", err)
		}
		log.Println("order id=>", data.OrderID, "Receiver=>", string(body))

	}()

}

func callbackhandler(c *gin.Context) {
	c.Data(200, "text", []byte(`success`))
}

type PolypayDepositeRequest struct {
	MerchantID string      `binding:"required" json:"merchant_id"`
	OrderID    string      `binding:"required" json:"order_id"`
	NotifyURL  string      `binding:"required" json:"notify_url"`
	Payment    string      `binding:"required" json:"payment"`
	BillPrice  string      `binding:"required" json:"bill_price"`
	Info       PolypayInfo `binding:"required" json:"info"`
}

type PolypayInfo struct {
	Telephone  string `binding:"required" json:"telephone"`
	PlayerID   string `binding:"required" json:"player_id"`
	Name       string `binding:"required" json:"name"`
	DeviceIP   string `binding:"required" json:"device_ip"`
	DeviceID   string `binding:"required" json:"device_id"`
	DeviceType string `binding:"required" json:"device_type"`
	PayAccount string `binding:"required" json:"pay_account"`
	Metadata   string `json:"metadata;omitempty"`
}

type SuccessResp struct {
	Status        int    `binding:"required" json:"status"`
	Message       string `binding:"required" json:"message"`
	TradeID       string `binding:"required" json:"trade_id"`
	OrderID       string `binding:"required" json:"order_id"`
	PayURL        string `binding:"required" json:"pay_url"`
	PayeeRealName string `binding:"required" json:"payee_realname"`
}

type CallBackReq struct {
	Status        int    `binding:"required" json:"status"`
	TradeAmount   string `binding:"required" json:"trade_amount"`
	ReceiptAmount string `binding:"required" json:"receipt_amount"`
	TradeID       string `binding:"required" json:"trade_id"`
	UpstreamOrder string `binding:"required" json:"upstream_order"`
	RepeatPay     bool   `binding:"required" json:"repeat_pay"`
	OrderID       string `binding:"required" json:"order_id"`
}
