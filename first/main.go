package main

import (
	"fmt"
	"github.com/DevOpserzhao/ops_gin/first/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/first/routers"
	"log"

	"github.com/DevOpserzhao/ops_gin/first/models"
	"github.com/DevOpserzhao/ops_gin/first/pkg/logging"
	"github.com/fvbock/endless"
	"syscall"
)

func main() {
	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

	setting.Setup()
	models.Setup()

	logging.Setup()
	//
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
