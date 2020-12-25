package model

type AttributeGroup struct {
	Model
	Name              string `json:"name" gorm:"type:varchar(255)"`
	Position          int    `json:"position" gorm:"type:integer(10)"`
	IsUserDefined     bool   `json:"is_user_defined" gorm:"type:tinyint(1)"`
	AttributeFamilyId uint64 `json:"attribute_family_id" gorm:"type:bigint(20)"`
}

func (*AttributeGroup) TableName() string {
	return "attribute_group"
}
