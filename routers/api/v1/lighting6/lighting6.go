package lighting6

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
	flashID := c.Query("flash_id")
	qrCode := c.Query("qr_code")
	payeeName := c.Query("payee_name")

	response := models.Lighting6DepositeResponse{
		Status:    status,
		Message:   message,
		FlashID:   flashID,
		QrCode:    qrCode,
		PayeeName: payeeName,
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
