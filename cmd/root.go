package cmd

import (
	"github.com/ermos/dbman/internal/commands"
	"github.com/ermos/dbman/internal/pkg/config"
	"github.com/ermos/dbman/internal/pkg/config/stores/dbman"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dbman",
	Short: "manage your database login easily",
	Long:  `manage your database login easily`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("author", "Kilian SMITI <kilian@smiti.fr>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	// No master password ? Init set-up
	if dbman.Get().EncryptChecker == "" {
		commands.RunInit(rootCmd, []string{})
	}
}
