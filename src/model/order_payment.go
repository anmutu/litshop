package model

type OrderPayment struct {
	Model
	OrderId    uint64 `json:"order_id"`
	CustomerId uint64 `json:"customer_id"`
}

func (*OrderPayment) TableName() string {
	return "order_payment"
}

func (*OrderPayment) Connection() string {
	return "default"
}
