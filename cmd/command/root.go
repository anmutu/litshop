package command

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "lit",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
