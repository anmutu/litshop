package migration

import "github.com/go-gormigrate/gormigrate/v2"

type Migration interface {
	Additionally() []*gormigrate.Migration
}

func GetMigrations() []*gormigrate.Migration {
	var migrations []*gormigrate.Migration
	migrations = make([]*gormigrate.Migration, 10)

	migrations = append(migrations, NewInitialMigration().Additionally()...)

	return migrations
}
