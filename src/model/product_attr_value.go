package model

type ProductAttrValue struct {
	Model
	ProductDataMode
	AttributeId uint64      `json:"attribute_id"`
	Value       interface{} `json:"value"`
	Desc        string      `json:"desc"`
}

func (*ProductAttrValue) TableName() string {
	return "product_attr_value"
}
