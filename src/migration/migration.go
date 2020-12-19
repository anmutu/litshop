package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"litshop/src/config/mysql"
)

type Migration interface {
	Additionally() []*gormigrate.Migration
}

func GetMigrations() map[string][]*gormigrate.Migration {
	var migrations map[string][]*gormigrate.Migration
	migrations = make(map[string][]*gormigrate.Migration, len(mysql.Connections()))

	for _, v := range mysql.Connections() {
		migrations[v] = append(migrations[v], NewInitialMigration().Additionally(v)...)
	}

	return migrations
}
