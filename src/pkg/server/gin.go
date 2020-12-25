package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"litshop/src/config"
	"litshop/src/lvm/runtime"
	"os"
	"os/signal"
)

var (
	r *gin.Engine
)

func setupGin() {
	gin.DisableConsoleColor()

	if config.GetString("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()
}

func signalListen() {
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, os.Kill)
	for {
		s := <-c
		fmt.Println("get signal:", s)

		// 结束
		runtime.ShutDown()

		os.Exit(0)
	}
}

func InitGinHttpServer() *gin.Engine {
	setupGin()
	go signalListen()
	return r
}
