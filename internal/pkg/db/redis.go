package db

import (
	"github.com/ermos/dbm/internal/pkg/term"
)

func RunLinuxRedis(c Config) error {
	var args []string

	term.RequireCommand("redis-cli")

	args = append(args, "-h")
	args = append(args, c.Host)

	if c.Port != "" {
		args = append(args, "-p")
		args = append(args, c.Port)
	}

	if c.Username != "" {
		args = append(args, "--user")
		args = append(args, c.Username)
	}

	if c.PlainPassword != "" {
		args = append(args, "--pass")
		args = append(args, c.PlainPassword)
	}

	return term.RunCommand("redis-cli", args, []string{})
}
