package commands

import (
	"fmt"
	"os/exec"
)

// Returns the type of a single command.
func Type(cmd string) (string, error) {
	if IsBuiltin(cmd) {
		return fmt.Sprintf("%s is a shell builtin", cmd), nil
	}

	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s is %s", cmd, path), nil
}
