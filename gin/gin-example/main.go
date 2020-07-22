package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/setting"
	"github.com/stephenchen/go-learning/gin/gin-example/routers"
	"log"
	"syscall"
)

func main() {
	// 1. kill -1 pid
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v\n", err)
	}

	// 2. http.Server - Shutdown()
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
