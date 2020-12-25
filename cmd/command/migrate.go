package command

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"litshop/src/config/mysql"
	"litshop/src/migration"
	"litshop/src/pkg/logger"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "执行数据库迁移",
	Run:   migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	migrations := migration.GetMigrations()
	fmt.Printf("migrations %#v \n", migrations)

	for conn, mList := range migrations {
		db := mysql.GormClientByConn(conn)
		m := gormigrate.New(db, gormigrate.DefaultOptions, mList)
		if err := m.Migrate(); err != nil {
			logger.Fatalf("Could not migrate: %v", err)
		}
	}

	logger.Printf("Migration did run successfully")
}
