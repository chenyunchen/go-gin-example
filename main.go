package main

import (
	"fmt"
	"log"
	"mockserver/pkg/logging"
	"mockserver/pkg/setting"
	"mockserver/pkg/util"
	"mockserver/routers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = setting.ServerSetting.HttpPort
	}
	endPoint := fmt.Sprintf(":%s", port)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
