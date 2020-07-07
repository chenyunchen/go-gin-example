package lighting3

import (
	"net/http"
	"strconv"

	"mockserver/models"
	"mockserver/pkg/logging"

	"github.com/gin-gonic/gin"
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
	isError, _ := strconv.ParseBool(strError)
	if isError {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// if success
	c.JSON(http.StatusOK, response)
	go DepositeCallback(c, status, tradeID, orderID)
}
