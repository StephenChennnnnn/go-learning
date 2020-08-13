package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/stephenchen/go-learning/gin/gin-example/models"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/logging"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/setting"
	"github.com/stephenchen/go-learning/gin/gin-example/routers"
	"log"
	"syscall"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()

}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	// 1.http.Server - Refactoring
	//gin.SetMode(setting.ServerSetting.RunMode)
	//
	//routersInit := routers.InitRouter()
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//maxHeaderBytes := 1 << 20
	//
	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler:        routersInit,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: maxHeaderBytes,
	//}
	//
	//log.Printf("[info] start http server listening %s", endPoint)
	//
	//server.ListenAndServe()

	// 2. kill -1 pid
	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
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
		log.Printf("Server err: %v\n", err)
	}

	// 3. http.Server - Shutdown()
	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:           router,
	//	ReadHeaderTimeout: setting.ReadTimeout,
	//	WriteTimeout:      setting.WriteTimeout,
	//	MaxHeaderBytes:    1 << 20,
	//}
	//
	//go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Print("Listen: %s\n", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//
	//log.Println("Shutdown Server...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown: ", err)
	//}
	//
	//log.Println("Server exiting")
}
