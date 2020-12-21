package model

import "golang.org/x/mobile/geom"

type Gender = int

type Address struct {
	Model
	CustomerId  uint64      `json:"customer_id"`
	Consignee   int         `json:"consignee"`
	AddressType int         `json:"address_type"`
	Name        string      `json:"name"`
	Gender      Gender      `json:"gender"`
	CompanyName string      `json:"company_name"`
	Address1    string      `json:"address1"`
	Address2    string      `json:"address2"`
	Postcode    string      `json:"postcode"`
	Street      string      `json:"street"`
	Area        string      `json:"area"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	AreaNo      string      `json:"area_no"`
	CityNo      string      `json:"city_no"`
	StateNo     string      `json:"state_no"`
	CountryNo   string      `json:"country_no"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Label       string      `json:"label"`
	Centre      geom.Point  `json:"centre" gorm:"type:point"`
	PostalCode  string      `json:"postal_code"`
	IsDefault   bool        `json:"is_default"`
	Additional  interface{} `json:"additional" gorm:"type:varchar(255)"`
}

func (*Address) TableName() string {
	return "attribute"
}
