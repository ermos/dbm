package commands

import (
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbm/internal/pkg/db"
	"github.com/ermos/dbm/internal/pkg/term"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"runtime"
)

func RunOpen(cmd *cobra.Command, args []string) {
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

	switch dbConfig.Protocol {
	case db.ProtocolMySQL:
		if runtime.GOOS == "windows" {
			panic("mysql protocol with windows isn't supported currently")
		} else {
			term.RequireCommand("mysql")
			if err = db.RunLinuxMySQL(dbConfig); err != nil {
				panic(err)
			}
		}
	default:
		panic("unsupported protocol")
	}
}
