package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"litshop/src/model"
)

type CustomerMigration struct {
}

func (*CustomerMigration) Additionally() []*gormigrate.Migration {
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
				return tx.AutoMigrate(&model.CustomerLogins{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable((&model.CustomerLogins{}).TableName())
			},
		},
	}
}

func NewCustomerMigration() *CustomerMigration {
	return new(CustomerMigration)
}
