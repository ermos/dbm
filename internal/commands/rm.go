package commands

import (
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Rm struct{}

func (Rm) Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		panic(
			"you must provide the alias name, for example: " +
				color.New(color.Bold).Sprint("dbm rm {alias}"),
		)
	}

	auth.WithMasterPassword()

	err := credentials.Get().RemoveAlias(args[0])
	if err != nil {
		panic(err)
	}
}
