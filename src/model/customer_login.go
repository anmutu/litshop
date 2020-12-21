package model

import "litshop/src/lvm/types"

type CustomerLogin struct {
	Model
	CustomerId uint64 `json:"customer_id"`

	AuthType    types.CustomerAuthType    `json:"auth_type"`
	AuthId      string                    `json:"auth_id"`
	AuthValue   string                    `json:"auth_value"`
	ExtraValue1 string                    `json:"extra_value_1"`
	ExtraValue2 string                    `json:"extra_value_2"`
	Status      types.CustomerLoginStatus `json:"status"`
}

func (*CustomerLogin) Connection() string {
	return "customer"
}

func (*CustomerLogin) TableName() string {
	return "customer_login"
}
