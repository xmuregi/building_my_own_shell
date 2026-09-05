package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xmuregi/building_my_own_shell/app/config"
	"github.com/xmuregi/building_my_own_shell/app/input"
)

// Allowed Shell builtins
var builtinList map[string]bool = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

// Returns bool value for Shell builtins
func IsBuiltin(command string) bool {
	return builtinList[command]
}

// Run external commands
func RunCommand(prompt *input.Prompt, binPaths *config.BinPath) {
	if IsBuiltin(prompt.Command) {
		switch prompt.Command {
		case "echo":
			Echo(prompt.Arg)
			return
		case "type":
			Type(prompt.Arg, binPaths)
			return
		}
	}
	path, err := binPaths.GetPath(prompt.Command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: not found\n", prompt.Command)
		return
	}
	if path != "" {
		argList := strings.Split(prompt.Arg, " ")

		cmd := exec.Command(prompt.Command, argList...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		_ = cmd.Run()
	}
}
