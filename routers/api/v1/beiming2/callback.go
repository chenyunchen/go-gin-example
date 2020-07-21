package beiming2

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
	targetAccount := c.Query("target_account")
	targetName := c.Query("target_name")
	orderFee := c.Query("order_fee")

	var payInfos []models.PayInfo
	payInfo := models.PayInfo{
		TargetAccount: targetAccount,
		TargetName:    targetName,
	}
	request := models.Beiming2DepositeCallbackRequest{
		MerchantCode:  merchantCode,
		MerchantOrder: merchantOrder,
		FlashID:       flashID,
		ReceiptPrice:  receiptPrice,
		PayedTime:     payedTime,
		PayInfo:       append(payInfos, payInfo),
		OrderFee:      orderFee,
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
