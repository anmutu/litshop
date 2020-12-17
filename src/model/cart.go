package model

import (
	"time"
)

type CartItemStatus = int

type Cart struct {
	Model
	CustomerDataMode
	ProductDataMode
	AddAt      time.Time      `json:"add_at"`
	SkuId      uint64         `json:"sku_id"`
	Quantity   uint64         `json:"quantity"`
	Status     CartItemStatus `json:"status"`
	Additional interface{}    `json:"additional"`
}

func (*Cart) TableName() string {
	return "cart"
}
