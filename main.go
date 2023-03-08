package main

import (
	"github.com/ermos/dbm/cmd"
	"github.com/ermos/dbm/internal/pkg/term"
)

func main() {
	defer term.ErrorHandler()
	cmd.Execute()
}
