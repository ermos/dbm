package cmd

import (
	"github.com/ermos/dbm/internal/commands"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a database",
	Long:  `Remove a database`,
	Run:   commands.RunRm,
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
