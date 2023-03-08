package commands

import (
	"fmt"
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores/dbm"
	"github.com/ermos/dbm/internal/pkg/term"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func RunInit(cmd *cobra.Command, args []string) {
	initText()
	fmt.Println("Your master password :")
	if err := auth.PromptMasterPassword(); err != nil {
		log.Fatal(err)
	}

	firstMasterPassword := auth.String()

	if len(firstMasterPassword) < 6 {
		fmt.Println(
			"Your master password must be greater than " +
				color.New(color.Bold).Sprint("6 characters") + ".",
		)
		os.Exit(0)
	}

	initText()
	fmt.Println("Verify your master password :")
	if err := auth.PromptMasterPassword(); err != nil {
		log.Fatal(err)
	}

	if firstMasterPassword != auth.String() {
		fmt.Println("The master passwords you entered do not match. Please try again.")
		os.Exit(0)
	}

	err := dbm.Get().GenerateEncryptChecker(firstMasterPassword)
	if err != nil {
		fmt.Println("A problem occurred. Please try again.")
		os.Exit(0)
	}
}

func initText() {
	term.Clear()
	fmt.Print("Welcome to ", color.New(color.Bold).Sprint("dbm"), "!\n",
		"Please define your master password before continuing.\n",
		"Use a strong and unique master password,\n",
		"it will be asking before each command.\n\n",
	)
}
