package cmd

import (
	"github.com/ermos/dbm/internal/commands"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Store a new database",
	Long:  `Store a new database`,
	Run:   commands.RunAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
