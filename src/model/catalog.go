package model

import "litshop/src/core/types"

// 店铺商品分类
type Catalog struct {
	Model
	Pid           uint64          `json:"pid"`
	Name          string          `json:"name"`
	Desc          string          `json:"desc"`
	Status        types.ComStatus `json:"status"`
	Path          string          `json:"path"`
	LogoMediaId   uint64          `json:"logo_media_id"`
	BannerMediaId uint64          `json:"banner_media_id"`
}

func (*Catalog) TableName() string {
	return "catalog"
}
