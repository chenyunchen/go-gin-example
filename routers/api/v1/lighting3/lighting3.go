package lighting3

import (
	"github.com/gin-gonic/gin"
	"mockserver/models"
	"mockserver/pkg/logging"
	"net/http"
	"strconv"
)

// response only handler
func ResponseOnly(c *gin.Context) {
	// required
	strStatus := c.Query("status")
	status, err := strconv.Atoi(strStatus)
	if err != nil {
		logging.Error(err)
	}
	message := c.Query("message")
	tradeID := c.Query("trade_id")
	orderID := c.Query("order_id")
	payURL := c.Query("pay_url")

	response := models.Lighting3DepositeResponse{
		Status:  status,
		Message: message,
		TradeID: tradeID,
		OrderID: orderID,
		PayURL:  payURL,
	}

	// if fail (option)
	strError := c.Query("error")
	isError, err := strconv.ParseBool(strError)
	if err != nil {
		logging.Error(err)
	}
	if isError {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// if success
	c.JSON(http.StatusOK, response)
}
