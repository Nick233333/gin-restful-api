package main

import (
	"fmt"
	"log"
	"syscall"

	"gin-restful-api/models"
	"gin-restful-api/pkg/logging"
	"gin-restful-api/pkg/setting"
	"gin-restful-api/routers"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
