package db

import (
	"github.com/ermos/dbm/internal/pkg/term"
)

func RunLinuxMongoDB(c Config) error {
	var args []string

	term.RequireCommand("mysql")

	args = append(args, "--host")
	args = append(args, c.Host)

	if c.Port != "" {
		args = append(args, "--port")
		args = append(args, c.Port)
	}

	if c.Username != "" {
		args = append(args, "-u")
		args = append(args, c.Username)
	}

	if c.PlainPassword != "" {
		args = append(args, "-p")
		args = append(args, c.PlainPassword)
	}

	if c.DefaultDatabase != "" {
		args = append(args, "--db")
		args = append(args, c.DefaultDatabase)
	}

	return term.RunCommand("mongosh", args, []string{})
}
