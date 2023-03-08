package main

import (
	"github.com/ermos/dbman/cmd"
	"github.com/ermos/dbman/internal/pkg/term"
)

func main() {
	defer term.ErrorHandler()
	cmd.Execute()
}
