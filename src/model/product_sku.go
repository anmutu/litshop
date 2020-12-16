package model

import (
	"golang.org/x/text/currency"
	"litshop/src/core/types"
)

type ProductSku struct {
	ProductDataMode
	AttributeList  []uint64        `json:"attribute_list"`
	BannerMediaId  uint64          `json:"banner_media_id"`
	MainPicMediaId uint64          `json:"main_pic_media_id"`
	Fee            currency.Amount `json:"fee"`
	MarketingFee   currency.Amount `json:"marketing_fee"`
	Status         types.ComStatus `json:"status"`
}

func (*ProductSku) TableName() string {
	return "product_sku"
}
