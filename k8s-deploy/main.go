package main

import (
	"flag"
	"fmt"
	clientgo "github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/client-go"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/logging"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func init() {

	env := flag.String("ENV_OPTS", "binary", "k8s or binary")
	flag.Parse()

	setting.Setup()
	logging.SetupAecss()
	clientgo.Setup(*env)
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger k8s运维平台

// @contact.name zxf
// @contact.url
// @contact.email

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

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
