package term

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func RunCommand(name string, args []string, envs []string) error {
	cmd := exec.Command(name, args...)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, envs...)
	// Disable history
	cmd.Env = append(cmd.Env, "HISTCONTROL=ignorespace")

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
