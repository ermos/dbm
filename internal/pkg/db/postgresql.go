package db

import (
	"fmt"
	"github.com/ermos/dbm/internal/pkg/term"
)

func RunLinuxPostgreSQL(c Config) error {
	var args []string
	var env []string

	term.RequireCommand("psql")

	args = append(args, fmt.Sprintf("--host=%s", c.Host))

	if c.Port != "" {
		args = append(args, fmt.Sprintf("--port=%s", c.Port))
	}

	if c.Username != "" {
		args = append(args, fmt.Sprintf("--username=%s", c.Username))
	}

	if c.DefaultDatabase != "" {
		args = append(args, c.DefaultDatabase)
	}

	if c.PlainPassword != "" {
		env = append(env, fmt.Sprintf("PGPASSWORD=%s", c.PlainPassword))
	}

	return term.RunCommand("psql", args, env)
}
