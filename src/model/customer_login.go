package model

import "litshop/src/lvm/types"

type CustomerLogin struct {
	Model
	CustomerId uint64 `json:"customer_id"`

	AuthType    types.CustomerAuthType    `json:"auth_type" gorm:"type:tinyint"`
	AuthId      string                    `json:"auth_id" gorm:"type:varchar(255)"`
	AuthValue   string                    `json:"auth_value" gorm:"type:varchar(255)"`
	ExtraValue1 string                    `json:"extra_value_1" gorm:"type:json"`
	ExtraValue2 string                    `json:"extra_value_2" gorm:"type:json"`
	Status      types.CustomerLoginStatus `json:"status" gorm:"type:tinyint"`
}

func (*CustomerLogin) Connection() string {
	return "customer"
}

func (*CustomerLogin) TableName() string {
	return "customer_login"
}
