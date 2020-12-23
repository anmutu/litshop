package model

import "golang.org/x/mobile/geom"

type Gender = int

//go:generate genrepository -dir . -structs=Address
type Address struct {
	Model
	CustomerId  uint64      `json:"customer_id"`
	Consignee   string      `json:"consignee" gorm:"type:varchar(255)"`
	AddressType int         `json:"address_type" gorm:"type:tinyint"`
	Name        string      `json:"name" gorm:"type:tinyint"`
	Gender      Gender      `json:"gender" gorm:"type:varchar(255)"`
	CompanyName string      `json:"company_name" gorm:"type:varchar(255)"`
	Address1    string      `json:"address1" gorm:"type:varchar(255)"`
	Address2    string      `json:"address2" gorm:"type:varchar(255)"`
	Postcode    string      `json:"postcode" gorm:"type:varchar(255)"`
	Street      string      `json:"street" gorm:"type:varchar(255)"`
	Area        string      `json:"area" gorm:"type:varchar(255)"`
	City        string      `json:"city" gorm:"type:varchar(255)"`
	State       string      `json:"state" gorm:"type:varchar(255)"`
	Country     string      `json:"country" gorm:"type:varchar(255)"`
	AreaNo      string      `json:"area_no" gorm:"type:bigint(20)"`
	CityNo      string      `json:"city_no" gorm:"type:bigint(20)"`
	StateNo     string      `json:"state_no" gorm:"type:bigint(20)"`
	CountryNo   string      `json:"country_no" gorm:"type:bigint(20)"`
	Email       string      `json:"email" gorm:"type:bigint(20)"`
	Phone       string      `json:"phone" gorm:"type:bigint(20)"`
	Label       string      `json:"label" gorm:"type:bigint(20)"`
	Centre      geom.Point  `json:"centre" gorm:"type:point"`
	PostalCode  string      `json:"postal_code" gorm:"type:bigint(20)"`
	IsDefault   bool        `json:"is_default"`
	Additional  interface{} `json:"additional" gorm:"type:varchar(255)"`
}

func (*Address) TableName() string {
	return "attribute"
}
