package util

import (
	"fmt"
	"os/exec"
)

func RunShell(shell string, args ...interface{}) error {
	shell = fmt.Sprintf(shell, args...)
	cmd := exec.Command("bash", "-c", shell)
	err := cmd.Run()
	return err
}
