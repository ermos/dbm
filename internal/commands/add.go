package commands

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbm/internal/pkg/db"
	"github.com/ermos/dbm/internal/pkg/term"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Add struct{}

func (Add) text() {
	term.Clear()
	term.TitlePrint("Adding a new database to dbm")
	fmt.Print("This command is designed to guide you through\n" +
		"the process of adding a new database to " + color.New(color.Bold).Sprint("dbm") + ".\n\n",
	)
}

func (c Add) Run(cmd *cobra.Command, args []string) {
	var dbConfig db.Config

	auth.WithMasterPassword()

	c.text()
	err := survey.Ask(c.getSurvey(), &dbConfig)
	if err != nil {
		panic(err)
	}

	if err = credentials.Get().Add(dbConfig); err != nil {
		panic(err)
	}

	if err = credentials.Get().Save(); err != nil {
		panic(err)
	}
}

func (c Add) getSurvey() []*survey.Question {
	return []*survey.Question{
		{
			Name:      "alias",
			Prompt:    &survey.Input{Message: "What is the alias name ?"},
			Validate:  c.requiredUniqueAlias,
			Transform: survey.ToLower,
		},
		{
			Name: "protocol",
			Prompt: &survey.Select{
				Message: "What is the protocol ?",
				Options: db.GetProtocols(),
				Default: "mysql",
			},
			Validate: survey.Required,
		},
		{
			Name:      "host",
			Prompt:    &survey.Input{Message: "What is the hostname ?"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "port",
			Prompt:    &survey.Input{Message: "What is the port ?"},
			Transform: survey.ToLower,
		},
		{
			Name:      "username",
			Prompt:    &survey.Input{Message: "What is your database username ? (not required)"},
			Transform: survey.ToLower,
		},
		{
			Name:   "plainPassword",
			Prompt: &survey.Password{Message: "What is your database password ? (not required)"},
		},
		{
			Name:   "defaultDatabase",
			Prompt: &survey.Input{Message: "What is the default database ? (not required)"},
		},
	}
}

func (Add) requiredUniqueAlias(val interface{}) error {
	err := survey.Required(val)
	if err != nil {
		return err
	}

	if credentials.Get().Credentials[fmt.Sprintf("%v", val)].Alias != "" {
		return errors.New("alias already used")
	}

	return nil
}
