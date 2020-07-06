package polypay

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"mockserver/models"
	"mockserver/pkg/logging"

	"github.com/gin-gonic/gin"
)

func DepositeCallback(c *gin.Context, tradeID, orderID string) {
	strStatus := c.Query("status")
	status, err := strconv.Atoi(strStatus)
	if err != nil {
		logging.Error(err)
		return
	}

	tradeAmount := c.Query("trade_amount")
	receiptAmount := c.Query("receipt_amount")
	upstreamOrder := c.Query("upstream_amount")
	strRepeatPay := c.Query("repeat_pay")
	repeatPay, err := strconv.ParseBool(strRepeatPay)
	if err != nil {
		logging.Error(err)
		return
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
		return
	}

	notifyURL := c.Query("notify_url")
	strRetry := c.Query("retry")
	retry, err := strconv.Atoi(strRetry)
	if err != nil {
		logging.Error(err)
	}

	for i := 1; i <= 1+retry; i++ {
		time.Sleep(5 * time.Second)
		_, err := http.Post(notifyURL, "application/json", bytes.NewBuffer(byteData))
		if err != nil {
			logging.Error(err)
			continue
		}
		break
	}
}
