package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/xmuregi/building_my_own_shell/app/config"
)

// Takes in commands and prints what type they are
func Type(arg string, binPaths *config.BinPath) {

	argList := strings.Split(arg, " ")

	for _, cmd := range argList {
		found := false

		// Incase its a shell builtin
		if IsCommand(cmd) {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
			found = true
			continue
		}

		// Incase its an external executable
		path, err := binPaths.GetPath(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
			continue
		}

		if path != ""{
			fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, path)
			found = true
			continue
		}

		// Incase its never found
		if !found {
			fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
		}
	}

}
