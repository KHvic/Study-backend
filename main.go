package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KHvic/quiz-backend/dao"
	"github.com/KHvic/quiz-backend/pkg/logging"
	"github.com/KHvic/quiz-backend/pkg/setting"
	"github.com/KHvic/quiz-backend/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	dao.Setup()
	logging.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
