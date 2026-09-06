package commands

import (
	"fmt"
	"os"
)

// cd (change directory) builtin is used to change the current working directory.
func Cd(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", path)
	}
	return nil
}
