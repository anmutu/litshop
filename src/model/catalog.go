package model

import "litshop/src/lvm/types"

// 店铺商品分类
type Catalog struct {
	Model
	Pid           uint64          `json:"pid"`
	Name          string          `json:"name" gorm:"type:varchar(255)"`
	Desc          string          `json:"desc" gorm:"type:varchar(255)"`
	Status        types.ComStatus `json:"status" gorm:"type:integer(10)"`
	Path          string          `json:"path" gorm:"type:varchar(255)"`
	LogoMediaId   uint64          `json:"logo_media_id"`
	BannerMediaId uint64          `json:"banner_media_id"`
}

func (*Catalog) TableName() string {
	return "catalog"
}
