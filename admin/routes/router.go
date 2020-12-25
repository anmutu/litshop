package routes

import (
	"github.com/gin-gonic/gin"
	"litshop/admin/controller"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("echo", controller.CtrlWrapper(controller.HealthCheck))

	authR := r.Group("/auth")
	{
		authRoutes(authR)
	}
}
