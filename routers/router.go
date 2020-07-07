package routers

import (
	"github.com/gin-gonic/gin"
	"mockserver/pkg/logging"
	"mockserver/routers/api/v1/bankcard"
	"mockserver/routers/api/v1/beiming1"
	"mockserver/routers/api/v1/beiming2"
	"mockserver/routers/api/v1/exchange"
	"mockserver/routers/api/v1/lighting3"
	"mockserver/routers/api/v1/lighting6"
	"mockserver/routers/api/v1/polypay"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", defaultHandler)
	r.POST("/", defaultHandler)

	depositeV1 := r.Group("/deposite")
	{
		depositeV1.GET("/polypay", polypay.ResponseOnly)
		depositeV1.GET("/lighting3", lighting3.ResponseOnly)
		depositeV1.GET("/lighting6", lighting6.ResponseOnly)
		depositeV1.GET("/beiming1", beiming1.ResponseOnly)
		depositeV1.GET("/beiming2", beiming2.ResponseOnly)

		// simulate deposite handler
		depositeV1.POST("/polypay", polypay.Deposite)
	}

	withdrawV1 := r.Group("/withdraw")
	{
		withdrawV1.GET("/bankcard", bankcard.ResponseOnly)
		withdrawV1.GET("/exchange", exchange.ResponseOnly)

	}

	return r
}

func defaultHandler(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		logging.Error(err)
		return
	}

	logging.Info(string(body))
	c.JSON(http.StatusOK, "success")
}
