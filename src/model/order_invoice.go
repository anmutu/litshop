package model

type OrderInvoice struct {
	Model
	OrderId    uint64 `json:"order_id"`
	CustomerId uint64 `json:"customer_id"`
}

func (*OrderInvoice) TableName() string {
	return "order_invoice"
}

func (*OrderInvoice) Connection() string {
	return "default"
}
