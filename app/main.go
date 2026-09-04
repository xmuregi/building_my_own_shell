package main

import (
	"fmt"
	"os"

	"github.com/xmuregi/building_my_own_shell/app/input"
)

func main() {
	fmt.Printf("$ ")
	command, err := input.GetInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read command: %v\n", err)
		return
	}
	if !input.IsCommand(command) {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
	}
}
