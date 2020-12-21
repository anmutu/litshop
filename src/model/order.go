package model

type Order struct {
	Model
}

func (*Order) TableName() string {
	return "order"
}

func (*Order) Connection() string {
	return "default"
}
