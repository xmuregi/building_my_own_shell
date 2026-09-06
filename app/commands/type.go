package commands

import (
	"fmt"
	"os/exec"
)

// Takes in commands and prints what type they are
func Type(arg string, binPaths *config.BinPath) {

	argList := strings.Split(arg, " ")

	for _, cmd := range argList {
		found := false
		// Incase its a shell builtin
		if IsBuiltin(cmd) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
			found = true
			continue
		}

	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s is %s", cmd, path), nil
}
