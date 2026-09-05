package commands

import (
	"fmt"
	"github.com/xmuregi/building_my_own_shell/app/config"
	"os"
	"slices"
	"strings"
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
		for _, path := range binPaths.Paths {
			entries := binPaths.Entries[path]
			if slices.Contains(entries, cmd) {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", cmd, path, cmd)
				found = true
				break
			}
		}

		// Incase its never found
		if !found {
			fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
		}
	}

}
