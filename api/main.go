package main

import (
	"litshop/api/routes"
	_ "litshop/src/bootstrap"
	"litshop/src/pkg/logger"
	"litshop/src/pkg/server"
)

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"litshop/src/config"
)

var (
	addr string
)

func main() {
	addr = fmt.Sprintf("%s:%d", config.GetString("api.host"), config.GetInt("api.port"))
	r := server.InitGinHttpServer()
	r.Use(gin.Logger())
	routes.ApiRoutes(r)
	logger.Info("app running on ", addr)
	_ = r.Run(addr)
}
