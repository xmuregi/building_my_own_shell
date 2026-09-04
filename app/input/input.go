package input

import (
	"bufio"
	"os"
	"strings"
)

var commandList map[string]bool = make(map[string]bool)

// Get User Input from Standard Input
func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(command), err
}

func IsCommand(command string) bool {
	return commandList[command]
}
