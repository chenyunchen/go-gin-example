package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Response setting gin.JSON
func (g *Gin) ErrResponse(httpCode, errCode int, message string) {
	g.C.JSON(httpCode, Response{
		Status:  errCode,
		Message: message,
	})
	return
}
