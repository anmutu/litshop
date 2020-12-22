package handler

import (
	"github.com/gin-gonic/gin"
	"litshop/src/pkg/d"
)

func EchoHandler(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"status":    "ok",
		"checklist": d.H{},
	})
}
