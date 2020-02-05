package main

import (
	"fmt"
	client_go "github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/client-go"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/logging"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func init() {
	setting.Setup()
	logging.Setup()
	client_go.Setup()
}
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
