package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"litshop/src/model"
)

// 初始化db
type InitialMigration struct {
}

func (*InitialMigration) Additionally(conn string) []*gormigrate.Migration {
	switch conn {
	case "customer":
		return []*gormigrate.Migration{
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Customer{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Customer{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.CustomerLogin{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.CustomerLogin{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.CustomerProfile{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.CustomerProfile{}).TableName())
				},
			},
		}

	default:
		return []*gormigrate.Migration{
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Address{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Address{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Attribute{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Attribute{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.AttributeFamily{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.AttributeFamily{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.AttributeGroup{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.AttributeGroup{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.AttributeGroupMapping{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.AttributeGroupMapping{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.AttributeOption{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.AttributeOption{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Brand{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Brand{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Cart{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Cart{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Catalog{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Catalog{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Category{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Category{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Media{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Media{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Order{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Order{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Product{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.Product{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.ProductAttrValue{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.ProductAttrValue{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.ProductCatalogMapping{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.ProductCatalogMapping{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.ProductImageMapping{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.ProductImageMapping{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.ProductSku{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.ProductSku{}).TableName())
				},
			},
			{
				ID: "20201210",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.ProductSkuStock{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable((&model.ProductSkuStock{}).TableName())
				},
			},
		}
	}
}

func NewInitialMigration() *InitialMigration {
	return new(InitialMigration)
}
