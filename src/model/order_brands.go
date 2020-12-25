package model

type OrderBrand struct {
	Model
	OrderId    uint64 `json:"order_id"`
	CustomerId uint64 `json:"customer_id"`
}

func (*OrderBrand) TableName() string {
	return "order_brand"
}

func (*OrderBrand) Connection() string {
	return "default"
}
