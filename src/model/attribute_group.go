package model

type AttributeGroup struct {
	Model
	Name              string `json:"name"`
	Position          int    `json:"position"`
	IsUserDefined     bool   `json:"is_user_defined"`
	AttributeFamilyId uint64 `json:"attribute_family_id"`
}

func (*AttributeGroup) TableName() string {
	return "attribute_group"
}
