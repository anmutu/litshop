package model

type OrderItem struct {
	Model
	OrderId    uint64 `json:"order_id"`
	CustomerId uint64 `json:"customer_id"`
}

func (*OrderItem) TableName() string {
	return "order_item"
}

func (*OrderItem) Connection() string {
	return "default"
}
