package model

import "litshop/src/lvm/types"

type Brand struct {
	Model
	Name        string          `json:"name" gorm:"type:varchar(255)"`
	Desc        string          `json:"desc" gorm:"type:varchar(255)"`
	LogoMediaId uint64          `json:"logo_media_id"`
	Status      types.ComStatus `json:"status" gorm:"type:integer(10)"`
}

func (*Brand) TableName() string {
	return "brand"
}
