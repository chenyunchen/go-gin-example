package routers

import (
	"net/http"

	"mockserver/pkg/logging"
	"mockserver/routers/api/v1/bankcard"
	"mockserver/routers/api/v1/beiming1"
	"mockserver/routers/api/v1/beiming2"
	"mockserver/routers/api/v1/exchange"
	"mockserver/routers/api/v1/polypay"
	"mockserver/routers/api/v1/shan3"
	"mockserver/routers/api/v1/shan6"

	"github.com/gin-gonic/gin"
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
		depositeV1.GET("/shan3", shan3.ResponseOnly)
		depositeV1.GET("/shan6", shan6.ResponseOnly)
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
