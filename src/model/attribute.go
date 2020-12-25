package model

import "litshop/src/lvm/types"

//go:generate genmodel -dir . -structs=Address
type Attribute struct {
	Model
	Code             string              `json:"code" gorm:"type:varchar(255)"`
	Type             types.AttributeType `json:"type" gorm:"type:integer(10)"`
	Validation       string              `json:"validation" gorm:"type:varchar(255)"`
	Position         int                 `json:"position" gorm:"type:integer(10)"`
	IsRequired       bool                `json:"is_required" gorm:"type:tinyint(1)"`
	IsUnique         bool                `json:"is_unique" gorm:"type:tinyint(1)"`
	IsSkuAttribute   bool                `json:"is_sku_attribute" gorm:"type:tinyint(1)"`
	IsFilterable     bool                `json:"is_filterable" gorm:"type:tinyint(1)"`
	IsConfigurable   bool                `json:"is_configurable" gorm:"type:tinyint(1)"`
	IsUserDefined    bool                `json:"is_user_defined" gorm:"type:tinyint(1)"`
	IsVisibleOnFront bool                `json:"is_visible_on_front" gorm:"type:tinyint(1)"`
	SwatchType       int                 `json:"swatch_type" gorm:"type:integer(10)"`
	UseInFlat        int                 `json:"use_in_flat" gorm:"type:integer(10)"`
	IsComparable     bool                `json:"is_comparable" gorm:"type:tinyint(1)"`
}

func (*Attribute) TableName() string {
	return "attribute"
}
