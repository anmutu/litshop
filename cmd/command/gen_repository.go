package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"litshop/src/pkg/logger"
)

func init() {
	rootCmd.AddCommand(genRepositoryCmd)
}

var genRepositoryCmd = &cobra.Command{
	Use:   "make:repository",
	Short: "执行数据库迁移",
	Run:   genRepository,
}

func genRepository(cmd *cobra.Command, args []string) {
	fmt.Printf("genRepository %#v \n", args)

	logger.Printf("code generate successfully")
}
