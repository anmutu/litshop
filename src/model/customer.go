package model

import (
	"litshop/src/core/types"
)

type Customer struct {
	Model

	NickName string               `json:"nick_name"`
	Avatar   string               `json:"avatar"`
	Status   types.CustomerStatus `json:"status"`
}

func (*Customer) TableName() string {
	return "customer"
}

func (*Customer) Phone() string {
	return ""
}

func (*Customer) Email() string {
	return ""
}

func (*Customer) Token() string {
	return ""
}
