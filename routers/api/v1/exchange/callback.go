package exchange

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"mockserver/models"
	"mockserver/pkg/logging"

	"github.com/gin-gonic/gin"
)

func WithdrawCallback(c *gin.Context) {
	delay, _ := strconv.Atoi(c.DefaultQuery("delay", "5"))
	orderID := c.Query("order_id")
	tradeNo := c.Query("trade_No")
	serviceCharge := c.Query("service_charge")
	orderAmout := c.Query("order_amout")
	payAccount := c.Query("pay_account")
	payTime := c.Query("pay_time")
	message := c.Query("message")
	orderFee := c.Query("order_fee")

	strCode := c.Query("code")
	code, err := strconv.Atoi(strCode)
	if err != nil {
		logging.Error(err)
	}

	request := models.ExchangeWithdrawCallbackRequest{
		OrderID:       orderID,
		TradeNo:       tradeNo,
		ServiceCharge: serviceCharge,
		OrderAmout:    orderAmout,
		PayAccount:    payAccount,
		Code:          code,
		PayTime:       payTime,
		Message:       message,
		OrderFee:      orderFee,
	}
	byteData, err := json.Marshal(request)
	if err != nil {
		logging.Error(err)
	}

	callbackURL := c.Query("callback_url")
	strRetry := c.Query("retry")
	retry, err := strconv.Atoi(strRetry)
	if err != nil {
		logging.Error(err)
	}

	v := url.Values{}
	v.Add("t", c.Query("t"))
	v.Add("sign", c.Query("sign"))
	u := callbackURL + "?" + v.Encode()

	for i := 1; i <= 1+retry; i++ {
		time.Sleep(time.Duration(delay) * time.Second)
		_, err := http.Post(u, "application/json", bytes.NewBuffer(byteData))
		if err != nil {
			logging.Error(err)
			continue
		}
		logging.Info("Callback POST: " + u)
		logging.Info("Callback BODY: " + string(byteData))
		break
	}
}
