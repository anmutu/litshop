package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Modeler interface {
	schema.Tabler
	Connection() string
}

type Model struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (*Model) Connection() string {
	return "default"
}

type CustomerDataMode struct {
	CustomerId uint64 `gorm:"primarykey" json:"customer_id"`
}

type ProductDataMode struct {
	ProductId uint64 `gorm:"primarykey" json:"product_id"`
}
