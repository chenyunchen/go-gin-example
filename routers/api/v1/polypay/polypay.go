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

// response only handler
func ResponseOnly(c *gin.Context) {
	appG := app.Gin{C: c}

	// required
	strStatus := c.Query("status")
	status, err := strconv.Atoi(strStatus)
	if err != nil {
		logging.Error(err)
	}
	message := c.Query("message")

	// if fail (option)
	strError := c.Query("error")
	isError, _ := strconv.ParseBool(strError)
	if isError {
		appG.ErrResponse(http.StatusBadRequest, status, message)
		return
	}

	// if success
	tradeID := c.Query("trade_id")
	orderID := c.Query("order_id")
	payURL := c.Query("pay_url")
	account := c.Query("account")
	name := c.Query("name")

	response := models.PolypayDepositeResponse{
		Status:  status,
		Message: message,
		TradeID: tradeID,
		OrderID: orderID,
		PayURL:  payURL,
		Account: account,
		Name:    name,
	}
	c.JSON(http.StatusOK, response)
	go DepositeCallback(c, status, tradeID, orderID)
}

// simulate deposite handler
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
