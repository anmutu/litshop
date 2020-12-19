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

	c.SetHeader(ginCtx.Request.Header)

	return c, nil
}

func FromGrpcServer() (*Context, error) {
	c := New()
	return c, nil
}
