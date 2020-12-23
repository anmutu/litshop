package model

import "gorm.io/gorm"

type CustomerRepository struct {
	BaseRepository

	DB *gorm.DB
}

func (b *CustomerRepository) init() {
	b.DB = b.GetDb((&Customer{}).Connection())
}

func (b *CustomerRepository) Where() {

}

func (b *CustomerRepository) Transaction(fn func(tx *gorm.DB) error) error {
	return b.DB.Transaction(fn)
}

func NewCustomerRepository() *CustomerRepository {
	r := new(CustomerRepository)
	r.init()
	return r
}
