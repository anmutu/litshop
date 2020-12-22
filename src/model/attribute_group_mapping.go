package model

type AttributeGroupMapping struct {
	Model
	AttributeId      uint64 `json:"attribute_id"`
	AttributeGroupId uint64 `json:"attribute_group_id"`
	Position         int    `json:"position" gorm:"type:integer(10)"`
}

func (*AttributeGroupMapping) TableName() string {
	return "attribute_group_mapping"
}
