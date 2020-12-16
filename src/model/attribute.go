package model

import "litshop/src/core/types"

type Attribute struct {
	Model
	Code             string              `json:"code"`
	Type             types.AttributeType `json:"type"`
	Validation       string              `json:"validation"`
	Position         int                 `json:"position"`
	IsRequired       bool                `json:"is_required"`
	IsUnique         bool                `json:"is_unique"`
	IsSkuAttribute   bool                `json:"is_sku_attribute"`
	IsFilterable     bool                `json:"is_filterable"`
	IsConfigurable   bool                `json:"is_configurable"`
	IsUserDefined    bool                `json:"is_user_defined"`
	IsVisibleOnFront bool                `json:"is_visible_on_front"`
	SwatchType       int                 `json:"swatch_type"`
	UseInFlat        int                 `json:"use_in_flat"`
	IsComparable     bool                `json:"is_comparable"`
}

func (*Attribute) TableName() string {
	return "attribute"
}
