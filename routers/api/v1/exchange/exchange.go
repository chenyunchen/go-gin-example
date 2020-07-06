package exchange

import (
	"net/http"
	"strconv"

	"mockserver/pkg/logging"

	"github.com/gin-gonic/gin"
)

// response only handler
func ResponseOnly(c *gin.Context) {
	// required
	body := c.Query("body")

	// if fail (option)
	strError := c.Query("error")
	isError, err := strconv.ParseBool(strError)
	if err != nil {
		logging.Error(err)
	}
	if isError {
		c.JSON(http.StatusBadRequest, body)
		return
	}

	// if success
	c.JSON(http.StatusOK, body)
}
