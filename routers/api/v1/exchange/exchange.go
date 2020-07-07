package exchange

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// response only handler
func ResponseOnly(c *gin.Context) {
	// required
	body := c.Query("body")

	// if fail (option)
	strError := c.Query("error")
	isError, _ := strconv.ParseBool(strError)
	if isError {
		c.JSON(http.StatusBadRequest, body)
		return
	}

	// if success
	c.JSON(http.StatusOK, body)
}
