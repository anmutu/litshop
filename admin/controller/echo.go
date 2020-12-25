package controller

import (
	"litshop/src/lvm/context"
	"litshop/src/pkg/d"
)

func HealthCheck(c *context.Context, request interface{}) (interface{}, error) {
	return map[string]interface{}{
		"status":    "ok",
		"checklist": d.H{},
	}, nil
}
