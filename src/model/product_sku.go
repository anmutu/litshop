package model

import (
	"litshop/src/lvm/types"
)

type ProductSku struct {
	Model
	ProductId      uint64          `json:"product_id"`
	AttributeList  []uint64        `json:"attribute_list" gorm:"type:json"`
	BannerMediaId  uint64          `json:"banner_media_id"`
	MainPicMediaId uint64          `json:"main_pic_media_id"`
	Fee            uint64          `json:"fee"`
	MarketingFee   uint64          `json:"marketing_fee"`
	Status         types.ComStatus `json:"status" gorm:"type:integer(10)"`
}

func (*ProductSku) TableName() string {
	return "product_sku"
}
