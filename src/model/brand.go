package model

import "litshop/src/core/types"

type Brand struct {
	Model
	Name        string          `json:"name"`
	Desc        string          `json:"desc"`
	LogoMediaId uint64          `json:"logo_media_id"`
	Status      types.ComStatus `json:"status"`
}

func (*Brand) TableName() string {
	return "brand"
}
