package model

type ProductAttrValue struct {
	Model
	ProductId   uint64      `json:"product_id"`
	AttributeId uint64      `json:"attribute_id"`
	Value       interface{} `json:"value" gorm:"type:varchar(255)"`
	Desc        string      `json:"desc" gorm:"type:varchar(255)"`
}

func (*ProductAttrValue) TableName() string {
	return "product_attr_value"
}
