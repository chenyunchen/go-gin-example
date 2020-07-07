package bankcard

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

func WithdrawCallback(c *gin.Context, status int, message, flashID string) {
	merchant := c.Query("merchant")
	payedMoney := c.Query("payed_money")
	merchantOrderID := c.Query("merchant_order_id")
	payedTime := c.Query("payed_time")

	request := models.BankcardWithdrawCallbackRequest{
		FlashID:         flashID,
		Merchant:        merchant,
		Status:          status,
		PayedMoney:      payedMoney,
		MerchantOrderID: merchantOrderID,
		PayedTime:       payedTime,
		Message:         message,
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
