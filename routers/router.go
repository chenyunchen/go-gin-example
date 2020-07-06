package routers

import (
	"github.com/gin-gonic/gin"
	"mockserver/routers/api/v1/polypay"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/deposite")
	{
		apiv1.GET("/polypay", polypay.Deposite)
		apiv1.POST("/polypay", polypay.Deposite)
	}

	return r
}
