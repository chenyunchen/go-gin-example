package lighting6

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

func DepositeCallback(c *gin.Context, status int, message, flashID string) {
	merchant := c.Query("merchant")
	payedMoney := c.Query("payed_money")
	upstreamOrder := c.Query("upstream_order")
	repeatPay := c.Query("repeat_pay")
	merchantOrderID := c.Query("merchant_order_id")

	request := models.Lighting6DepositeCallbackRequest{
		FlashID:         flashID,
		Merchant:        merchant,
		Status:          status,
		PayedMoney:      payedMoney,
		UpstreamOrder:   upstreamOrder,
		RepeatPay:       repeatPay,
		MerchantOrderID: merchantOrderID,
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
