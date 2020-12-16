package model

type ProductSkuStock struct {
	Model
	ProductDataMode
	SkuId uint64 `json:"sku_id"`
	Stock uint64 `json:"stock"`
}

func (*ProductSkuStock) TableName() string {
	return "product_sku_stock"
}
