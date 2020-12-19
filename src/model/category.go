package model

import "litshop/src/lvm/types"

// 商品类目
type Category struct {
	Model
	Pid    uint64          `json:"pid"`
	Name   string          `json:"name"`
	Desc   string          `json:"desc"`
	Status types.ComStatus `json:"status"`
}

func (*Category) TableName() string {
	return "category"
}
