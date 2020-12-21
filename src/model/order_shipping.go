package model

type OrderShipping struct {
	Model
	OrderId    uint64 `json:"order_id"`
	CustomerId uint64 `json:"customer_id"`
}

func (*OrderShipping) TableName() string {
	return "order_shipping"
}

func (*OrderShipping) Connection() string {
	return "default"
}
