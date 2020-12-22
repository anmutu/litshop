package model

import "litshop/src/lvm/types"

// 商品类目
type Category struct {
	Model
	Pid    uint64          `json:"pid"`
	Name   string          `json:"name" gorm:"type:varchar(255)"`
	Desc   string          `json:"desc" gorm:"type:varchar(255)"`
	Status types.ComStatus `json:"status" gorm:"type:integer(10)"`
}

func (*Category) TableName() string {
	return "category"
}
