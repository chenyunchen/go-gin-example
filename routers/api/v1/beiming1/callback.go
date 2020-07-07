package beiming1

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

func DepositeCallback(c *gin.Context, merchantCode, merchantOrder, flashID string) {
	receiptPrice := c.Query("receipt_price")
	payedTime := c.Query("payed_time")

	orderTypeStr := c.Query("order_type")
	orderType, err := strconv.Atoi(orderTypeStr)
	if err != nil {
		logging.Error(err)
	}
	orderModeStr := c.Query("order_mode")
	orderMode, err := strconv.Atoi(orderModeStr)
	if err != nil {
		logging.Error(err)
	}
	userLevelStr := c.Query("user_level")
	userLevel, err := strconv.Atoi(userLevelStr)
	if err != nil {
		logging.Error(err)
	}

	targetAccount := c.Query("target_account")
	targetName := c.Query("target_name")
	tradeTime := c.Query("trade_time")
	repeatPay := c.Query("repeat_pay")

	request := models.Beiming1DepositeCallbackRequest{
		MerchantCode:  merchantCode,
		MerchantOrder: merchantOrder,
		FlashID:       flashID,
		ReceiptPrice:  receiptPrice,
		PayedTime:     payedTime,
		OrderType:     orderType,
		OrderMode:     orderMode,
		UserLevel:     userLevel,
		TargetAccount: targetAccount,
		TargetName:    targetName,
		TradeTime:     tradeTime,
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
