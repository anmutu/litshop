package model

type ProductImageMapping struct {
	Model
	ProductDataMode
	MediaId uint64 `json:"media_id"`
	IsMain  bool   `json:"is_main"`
}

func (*ProductImageMapping) TableName() string {
	return "product_image_mapping"
}
