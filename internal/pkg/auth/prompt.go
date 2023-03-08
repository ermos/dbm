package auth

import (
	"fmt"
	"github.com/ermos/dbm/internal/pkg/config/stores/dbm"
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

	fmt.Print("[dbm] master password: ")
	err := PromptMasterPassword()
	fmt.Print("\n")
	if err != nil {
		log.Fatal(err)
	}

	if !dbm.Get().IsValidMasterPassword(String()) {
		if try <= 1 {
			panic("3 incorrect master password attempts")
		}

		fmt.Println("Master password not valid, try again.")
		WithMasterPassword(try - 1)
	}
}
