package model

type ProductCatalogMapping struct {
	Model
	ProductDataMode
	CatalogId uint64 `json:"catalog_id"`
}

func (*ProductCatalogMapping) TableName() string {
	return "product_catalog_mapping"
}
