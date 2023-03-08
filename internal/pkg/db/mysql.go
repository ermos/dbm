package db

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func RunLinuxMySQL(c Config) error {
	var args []string

	args = append(args, "-h")
	args = append(args, c.Host)

	if c.Port != "" {
		args = append(args, "-P")
		args = append(args, c.Port)
	}

	args = append(args, "-u")
	args = append(args, c.Username)

	if c.PlainPassword != "" {
		args = append(args, "-p"+c.PlainPassword)
	}

	if c.DefaultDatabase != "" {
		args = append(args, c.DefaultDatabase)
	}

	fmt.Println("mysql", strings.Join(args, " "))

	cmd := exec.Command("mysql", args...)

	// Disable history
	cmd.Env = append(os.Environ(), "HISTCONTROL=ignorespace")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	errchan := make(chan error, 1)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGINT)

	go func() {
		errchan <- cmd.Run()
	}()

	select {
	case err := <-errchan:
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
					return fmt.Errorf("(mysql) exited with status %d", status.ExitStatus())
				}
			}
			return err
		}
	case <-sigchan:
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	}

	return nil
}
