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
	if err != nil {
		log.Fatalf("Error loading binary paths: %v\n", err)
	}

REPL:
	for {
		fmt.Printf("$ ")
		prompt, err := input.GetPrompt()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to read prompt: %v\n", err)
			return
		}
		switch prompt.Command {
		case "exit":
			fmt.Fprintf(os.Stdout, "Exiting shell!\n")
			break REPL
		}
		commands.RunCommand(prompt, binPath)
	}
}
