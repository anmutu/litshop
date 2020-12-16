package model

type AttributeOption struct {
	Model
	SortOrder   int    `json:"sort_order"`
	AttributeId uint64 `json:"attribute_id"`
	SwatchValue uint   `json:"swatch_value"`
}

func (*AttributeOption) TableName() string {
	return "attribute_option"
}
