package model

import "litshop/src/lvm/types"

type AttributeFamily struct {
	Model
	Code          int             `json:"code" gorm:"type:varchar(255)"`
	Name          int             `json:"name" gorm:"type:varchar(255)"`
	Status        types.ComStatus `json:"status" gorm:"type:integer(10)"`
	IsUserDefined bool            `json:"is_user_defined" gorm:"type:tinyint(1)"`
}

func (*AttributeFamily) TableName() string {
	return "attribute_family"
}
