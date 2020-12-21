package model

import (
	"litshop/src/lvm/types"
)

type Product struct {
	Model
	BrandId      uint64            `json:"brand_id"`
	CategoryId   uint64            `json:"category_id"`
	Name         string            `json:"name"`
	Desc         string            `json:"desc"`
	SellingPoint string            `json:"selling_point"`
	Unit         string            `json:"unit"`
	Fee          uint64            `json:"fee"`
	MarketingFee uint64            `json:"marketing_fee"`
	Status       types.ComStatus   `json:"status"`
	OnSale       bool              `json:"on_sale"`
	Type         types.ProductType `json:"type"`
}

func (*Product) TableName() string {
	return "product"
}
