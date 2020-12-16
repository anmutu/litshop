package routes

import (
	"github.com/gin-gonic/gin"
	"litshop/src/app/api/handler"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("echo", handler.EchoHandler)
}
