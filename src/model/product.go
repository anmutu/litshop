package model

import (
	"golang.org/x/text/currency"
	"litshop/src/core/types"
)

type Product struct {
	Model
	BrandId      uint64            `json:"brand_id"`
	CategoryId   uint64            `json:"category_id"`
	Name         string            `json:"name"`
	Desc         string            `json:"desc"`
	SellingPoint string            `json:"selling_point"`
	Unit         string            `json:"unit"`
	Fee          currency.Amount   `json:"fee"`
	MarketingFee currency.Amount   `json:"marketing_fee"`
	Status       types.ComStatus   `json:"status"`
	OnSale       bool              `json:"on_sale"`
	Type         types.ProductType `json:"type"`
}

func (*Product) TableName() string {
	return "product"
}
