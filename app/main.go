package main

import (
	"fmt"
	"os"

	"github.com/xmuregi/building_my_own_shell/app/commands"
	"github.com/xmuregi/building_my_own_shell/app/input"
)

func main() {

REPL:
	for {
		fmt.Printf("$ ")
		prompt, err := input.GetPrompt()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to read prompt: %v\n", err)
			return
		}
		IsCommandVal, err := commands.IsCommand(prompt.Command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to read prompt: %v\n", err)
			return
		}
		if !IsCommandVal {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", prompt.Command)
			continue
		}

		switch prompt.Command {
		case "exit":
			fmt.Println("Exiting shell!")
			break REPL
		case "echo":
			commands.Echo(prompt.Arg)
		}
	}

}
