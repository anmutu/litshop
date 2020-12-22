package model

import (
	"litshop/src/lvm/types"
)

type Product struct {
	Model
	BrandId      uint64            `json:"brand_id"`
	CategoryId   uint64            `json:"category_id"`
	Name         string            `json:"name" gorm:"type:varchar(255)"`
	Desc         string            `json:"desc" gorm:"type:varchar(255)"`
	SellingPoint string            `json:"selling_point" gorm:"type:varchar(255)"`
	Unit         string            `json:"unit" gorm:"type:varchar(255)"`
	Fee          uint64            `json:"fee"`
	MarketingFee uint64            `json:"marketing_fee"`
	Status       types.ComStatus   `json:"status" gorm:"type:integer(10)"`
	OnSale       bool              `json:"on_sale"`
	Type         types.ProductType `json:"type" gorm:"type:integer(10)"`
}

func (*Product) TableName() string {
	return "product"
}
