package model

type ProductCatalogMapping struct {
	Model
	ProductId uint64 `json:"product_id"`
	CatalogId uint64 `json:"catalog_id"`
}

func (*ProductCatalogMapping) TableName() string {
	return "product_catalog_mapping"
}
