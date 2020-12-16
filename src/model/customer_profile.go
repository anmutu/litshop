package model

type CustomerProfile struct {
	Model
	CustomerDataMode
}

func (*CustomerProfile) TableName() string {
	return "customer_profile"
}

func (*CustomerProfile) Connection() string {
	return "customer"
}
