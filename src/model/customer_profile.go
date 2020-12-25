package model

//go:generate genmodel -dir . -structs=CustomerProfile
type CustomerProfile struct {
	Model
	CustomerId uint64 `json:"customer_id"`
}

func (*CustomerProfile) TableName() string {
	return "customer_profile"
}

func (*CustomerProfile) Connection() string {
	return "customer"
}
