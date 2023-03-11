package commands

import (
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbm/internal/pkg/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Open struct{}

func (Open) Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		panic(
			"you must provide the alias name, for example: " +
				color.New(color.Bold).Sprint("dbm open {alias}"),
		)
	}

	auth.WithMasterPassword()

	dbConfig, err := credentials.Get().Get(args[0])
	if err != nil {
		panic(err)
	}

	err = credentials.Get().UpdateLastConnection(dbConfig.Alias)
	if err != nil {
		panic(err)
	}

	err = db.Run(dbConfig)
	if err != nil {
		panic(err)
	}
}
