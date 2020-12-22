package main

import (
	"litshop/admin/routes"
	_ "litshop/src/bootstrap"
	"litshop/src/lvm/runtime"
	"litshop/src/pkg/logger"
)

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"litshop/src/config"
	"os"
	"os/signal"
)

var (
	host string
	r    *gin.Engine
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

func main() {
	host = fmt.Sprintf(":%d", config.GetInt("port"))
	setupGin()
	r.Use(gin.Logger())
	routes.ApiRoutes(r)
	go signalListen()
	logger.Info("app running on ", host)
	_ = r.Run(host)
}
