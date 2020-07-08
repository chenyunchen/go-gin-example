package shan3

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

func DepositeCallback(c *gin.Context, status int, tradeID, orderID string) {
	tid := c.Query("tid")
	tradeAmount := c.Query("trade_amount")
	receiptAmount := c.Query("receipt_amount")
	userData := c.Query("user_data")

	request := models.Shan3DepositeCallbackRequest{
		Tid:           tid,
		Status:        status,
		TradeAmount:   tradeAmount,
		ReceiptAmount: receiptAmount,
		TradeID:       tradeID,
		OrderID:       orderID,
		UserData:      userData,
	}
	byteData, err := json.Marshal(request)
	if err != nil {
		logging.Error(err)
	}

	notifyURL := c.Query("notify_url")
	strRetry := c.Query("retry")
	retry, err := strconv.Atoi(strRetry)
	if err != nil {
		logging.Error(err)
	}

	v := url.Values{}
	v.Add("time", c.Query("time"))
	v.Add("sign", c.Query("sign"))
	u := notifyURL + "?" + v.Encode()

	for i := 1; i <= 1+retry; i++ {
		time.Sleep(5 * time.Second)
		_, err := http.Post(u, "application/json", bytes.NewBuffer(byteData))
		if err != nil {
			logging.Error(err)
			continue
		}
		break
	}
}
