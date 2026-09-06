package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xmuregi/building_my_own_shell/app/input"
)

// Allowed Shell builtins
var builtinList map[string]bool = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
}

// Returns bool value for Shell builtins
func IsBuiltin(command string) bool {
	return builtinList[command]
}

// Runs a built-in command using the provided output files.
func runBuiltin(command *input.Prompt, stdout, stderr *os.File) {
	switch command.Command {
	case "echo":
		result, err := Echo(command.Arg)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return
		}
		fmt.Fprintln(stdout, result)
		return
	case "type":
		for _, name := range strings.Split(command.Arg, " ") {
			result, err := Type(name)
			if err != nil {
				fmt.Fprintln(stderr, err)
				continue
			}
			fmt.Fprintln(stdout, result)
		}
		return
	case "pwd":
		result, err := Pwd()
		if err != nil {
			fmt.Fprintln(stderr, err)
			return
		}
		fmt.Fprintln(stdout, result)
		return

	}
}

// Runs an external command using the provided output files.
func runExternal(command *input.Prompt, stdout, stderr *os.File, path string) {
	argList := strings.Split(command.Arg, " ")

	cmd := exec.Command(path, argList...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(stderr, err)
	}
}

// Runs built-in or external commands.
func RunCommand(prompt *input.Prompt) {
	if IsBuiltin(prompt.Command) {
		runBuiltin(prompt, os.Stdout, os.Stderr)
		return
	}
	path, err := exec.LookPath(prompt.Command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: not found\n", prompt.Command)
		return
	}
	if path != "" {
		runExternal(prompt, os.Stdout, os.Stderr, path)
	}
}
