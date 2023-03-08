package cmd

import (
	"github.com/ermos/dbm/internal/commands"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a database connection",
	Long:  `Open a database connection`,
	Run:   commands.RunOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)
}
