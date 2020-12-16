package model

type Order struct {
}

func (*Order) TableName() string {
	return "order"
}

func (*Order) Connection() string {
	return "default"
}
