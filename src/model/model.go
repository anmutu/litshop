package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CustomerDataMode struct {
	CustomerId uint64 `gorm:"primarykey" json:"customer_id"`
}

type ProductDataMode struct {
	ProductId uint64 `gorm:"primarykey" json:"product_id"`
}
