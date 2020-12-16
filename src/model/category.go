package model

import "litshop/src/core/types"

// 商品类目
type Category struct {
	Model
	Pid    uint64          `json:"pid"`
	Name   string          `json:"name"`
	Desc   string          `json:"desc"`
	Status types.ComStatus `json:"status"`
}
