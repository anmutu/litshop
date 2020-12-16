package model

import "litshop/src/core/types"

type CustomerLogins struct {
	Model
	CustomerDataMode

	AuthType    types.CustomerAuthType    `json:"auth_type"`
	AuthId      string                    `json:"auth_id"`
	AuthValue   string                    `json:"auth_value"`
	ExtraValue1 string                    `json:"extra_value_1"`
	ExtraValue2 string                    `json:"extra_value_2"`
	Status      types.CustomerLoginStatus `json:"status"`
}

func (*CustomerLogins) TableName() string {
	return "customer_login"
}
