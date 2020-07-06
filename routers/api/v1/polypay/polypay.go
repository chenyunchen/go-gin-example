package polypay

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"mockserver/models"
	"mockserver/pkg/app"
	"mockserver/pkg/e"
	"mockserver/pkg/logging"
	"mockserver/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func Deposite(c *gin.Context) {
	appG := app.Gin{C: c}

	t := c.Query("time")
	tt, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		logging.Error(err)
		return
	}
	tm := time.Unix(tt, 0)
	if time.Since(tm).Seconds() > 150 {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		return
	}

	s := c.Query("sign")
	body, err := c.GetRawData()
	if err != nil {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		logging.Error(err)
		return
	}
	if s != sign(body, tm, setting.PaymentSetting.PolypayToken) {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		return
	}

	o := c.Query("origin")
	if o != setting.PaymentSetting.PolypayOrigin {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		return
	}

	os := c.Query("origin_sign")
	if os != sign(body, tm, setting.PaymentSetting.PolypayOriginToken) {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		return
	}

	request := models.PolypayDepositeRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		appG.ErrResponse(http.StatusBadRequest, e.POLYPAY_ERROR, e.GetMessage(e.INVALID_PARAMS))
		logging.Error(err)
		return
	}

	response := models.PolypayDepositeResponse{
		Status:  e.POLYPAY_SUCCESS,
		Message: e.GetMessage(e.POLYPAY_SUCCESS),
		TradeID: xid.New().String(),
		OrderID: request.OrderID,
		PayURL:  "https://t.cn/pay",
		Account: request.Info.PayAccount,
		Name:    request.Info.Name,
	}
	c.JSON(http.StatusOK, response)
}
