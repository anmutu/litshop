package model

type ProductImageMapping struct {
	Model
	ProductId uint64 `json:"product_id"`
	MediaId   uint64 `json:"media_id"`
	IsMain    bool   `json:"is_main"`
}

func (*ProductImageMapping) TableName() string {
	return "product_image_mapping"
}
