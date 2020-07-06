package beiming1

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
	merchantCode := c.Query("merchant_code")
	merchantOrder := c.Query("merchant_order")
	flashID := c.Query("flash_id")
	payURL := c.Query("pay_url")
	billPrice := c.Query("bill_price")
	cardNum := c.Query("card_num")
	acceptNum := c.Query("accept_num")
	bankName := c.Query("bank_name")
	province := c.Query("province")
	city := c.Query("city")
	country := c.Query("country")
	node := c.Query("node")
	createdAT := c.Query("created_at")
	remark := c.Query("remark")

	var data = models.Beiming1DepositeData{
		MerchantCode:  merchantCode,
		MerchantOrder: merchantOrder,
		FlashID:       flashID,
		PayUrl:        payURL,
		Status:        status,
		Message:       message,
		BillPrice:     billPrice,
		CardNum:       cardNum,
		AcceptName:    acceptNum,
		BankName:      bankName,
		Province:      province,
		City:          city,
		Country:       country,
		Node:          node,
		CreatedAt:     createdAT,
		Remark:        remark,
	}
	response := models.Beiming1DepositeResponse{
		Data: data,
	}

	// if fail (option)
	strError := c.Query("error")
	isError, err := strconv.ParseBool(strError)
	if err != nil {
		logging.Error(err)
	}
	if isError {
		response.Code = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// if success
	response.Code = http.StatusOK
	c.JSON(http.StatusOK, response)
}
