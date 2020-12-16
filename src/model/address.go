package model

type Gender = int

type Address struct {
	Model
	CustomerDataMode
	AddressType      int         `json:"address_type"`
	Name             string      `json:"name"`
	Gender           Gender      `json:"gender"`
	CompanyName      string      `json:"company_name"`
	Address1         string      `json:"address1"`
	Address2         string      `json:"address2"`
	Postcode         string      `json:"postcode"`
	Street           string      `json:"street"`
	Area             string      `json:"area"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	Country          string      `json:"country"`
	AreaNo           string      `json:"area_no"`
	CityNo           string      `json:"city_no"`
	StateNo          string      `json:"state_no"`
	CountryNo        string      `json:"country_no"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	IsDefaultAddress bool        `json:"is_default_address"`
	Additional       interface{} `json:"additional"`
}

func (*Address) TableName() string {
	return "attribute"
}
