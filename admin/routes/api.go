package routes

import (
	"github.com/gin-gonic/gin"
	"litshop/admin/handler"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("echo", handler.EchoHandler)
}
