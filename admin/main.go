package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"litshop/admin/routes"
	_ "litshop/src/bootstrap"
	"litshop/src/config"
	"litshop/src/lvm/runtime"
	"litshop/src/pkg/logger"
	"litshop/src/pkg/server"
)

func init() {
	runtime.SetRunApp("admin")
}

var (
	addr string
)

func main() {
	addr = fmt.Sprintf("%s:%d", config.GetString("admin.host"), config.GetInt("admin.port"))

	r := server.InitGinHttpServer()
	r.Use(gin.Logger())

	routes.ApiRoutes(r)

	logger.Info("app running on ", addr)
	_ = r.Run(addr)
}
