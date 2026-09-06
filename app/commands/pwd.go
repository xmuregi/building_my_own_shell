package commands

import (
	"os"
)

// Returns the working directory
func Pwd() (string, error) {
	wkdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wkdir, nil
}
