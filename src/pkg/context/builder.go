package context

import (
	"github.com/gin-gonic/gin"
)

var (
	ginCtx *gin.Context
)

func FromGinContext(ctx *gin.Context) (*Context, error) {
	c := New()

	ginCtx = ctx

	return c, nil
}
