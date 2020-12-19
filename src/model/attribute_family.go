package model

import "litshop/src/lvm/types"

type AttributeFamily struct {
	Model
	Code          int             `json:"code"`
	Name          int             `json:"name"`
	Status        types.ComStatus `json:"status"`
	IsUserDefined bool            `json:"is_user_defined"`
}

func (*AttributeFamily) TableName() string {
	return "attribute_family"
}
