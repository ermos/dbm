package db

import (
	"github.com/ermos/dbm/internal/pkg/term"
)

func RunLinuxMySQL(c Config) error {
	var args []string

	term.RequireCommand("mysql")

	args = append(args, "-h")
	args = append(args, c.Host)

	if c.Port != "" {
		args = append(args, "-P")
		args = append(args, c.Port)
	}

	if c.Username != "" {
		args = append(args, "-u")
		args = append(args, c.Username)
	}

	if c.PlainPassword != "" {
		args = append(args, "-p"+c.PlainPassword)
	}

	if c.DefaultDatabase != "" {
		args = append(args, c.DefaultDatabase)
	}

	return term.RunCommand("mysql", args, []string{})
}
