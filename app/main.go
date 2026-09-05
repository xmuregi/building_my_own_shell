package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xmuregi/building_my_own_shell/app/commands"
	"github.com/xmuregi/building_my_own_shell/app/config"
	"github.com/xmuregi/building_my_own_shell/app/input"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	binPath, err := config.NewBinPath()
	if err != nil{
		log.Fatalf("Error loading binary paths", err)
	}

REPL:
	for {
		fmt.Printf("$ ")
		prompt, err := input.GetPrompt()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to read prompt: %v\n", err)
			return
		}
		if !commands.IsCommand(prompt.Command) {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", prompt.Command)
			continue
		}

		switch prompt.Command {
		case "exit":
			fmt.Println("Exiting shell!")
			break REPL
		case "echo":
			commands.Echo(prompt.Arg)
		case "type":
			commands.Type(prompt.Arg, binPath)
		}
	}
}
