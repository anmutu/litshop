package controller

import (
	"github.com/gin-gonic/gin"
	"litshop/src/lvm/context"
	"litshop/src/pkg/d"
)

func CtrlWrapper(controller func(*context.Context, interface{}) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, err := context.FromGinContext(c)
		if err != nil {
			c.JSON(500, "")
			return
		}

		result, err := controller(ctx, d.H{})
		if err != nil {
			c.JSON(500, "")
			return
		}

		c.JSON(200, result)
		return
	}
}
