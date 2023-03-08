package cmd

import (
	"github.com/ermos/dbm/internal/commands"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show available databases",
	Long:  `Show available databases`,
	Run:   commands.RunList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
