package model

import (
	"time"
)

type CartItemStatus = int

type Cart struct {
	Model
	CustomerId uint64         `json:"customer_id"`
	ProductId  uint64         `json:"product_id"`
	AddAt      time.Time      `json:"add_at"`
	SkuId      uint64         `json:"sku_id"`
	Quantity   uint64         `json:"quantity"`
	Status     CartItemStatus `json:"status" gorm:"type:integer(10)"`
	Additional interface{}    `json:"additional" gorm:"type:varchar(255)"`
}

func (*Cart) TableName() string {
	return "cart"
}
