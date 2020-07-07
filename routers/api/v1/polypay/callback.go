package polypay

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
	tradeAmount := c.Query("trade_amount")
	receiptAmount := c.Query("receipt_amount")
	upstreamOrder := c.Query("upstream_amount")
	strRepeatPay := c.Query("repeat_pay")
	repeatPay, err := strconv.ParseBool(strRepeatPay)
	if err != nil {
		logging.Error(err)
	}

	request := models.PolypayDepositeCallbackRequest{
		TradeID:       tradeID,
		OrderID:       orderID,
		Status:        status,
		TradeAmount:   tradeAmount,
		ReceiptAmount: receiptAmount,
		UpstreamOrder: upstreamOrder,
		RepeatPay:     repeatPay,
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
