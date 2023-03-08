package auth

import (
	"fmt"
	"github.com/ermos/dbman/internal/pkg/config/stores/dbman"
	"golang.org/x/term"
	"log"
	"os"
)

var masterPassword []byte

func PromptMasterPassword() (err error) {
	masterPassword, err = term.ReadPassword(int(os.Stdin.Fd()))
	return
}

func String() string {
	return string(masterPassword)
}

func Bytes() []byte {
	return masterPassword
}

func WithMasterPassword(tryCount ...int) {
	try := 3
	if len(tryCount) != 0 {
		try = tryCount[0]
	}

	fmt.Print("[dbman] master password: ")
	err := PromptMasterPassword()
	fmt.Print("\n")
	if err != nil {
		log.Fatal(err)
	}

	if !dbman.Get().IsValidMasterPassword(String()) {
		if try <= 1 {
			panic("3 incorrect master password attempts")
		}

		fmt.Println("Master password not valid, try again.")
		WithMasterPassword(try - 1)
	}
}
