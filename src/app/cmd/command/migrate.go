package command

import (
	"github.com/spf13/cobra"
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

}
