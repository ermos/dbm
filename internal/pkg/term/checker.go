package term

import (
	"fmt"
	"os/exec"
)

func IsCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func RequireCommand(cmd string) {
	if !IsCommandExists(cmd) {
		panic(fmt.Sprintf("unsupported command %s", cmd))
	}
}
