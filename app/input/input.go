package input

import (
	"bufio"
	"os"
	"strings"
)

type Prompt struct {
	Command string
	Arg     string
}

// Get User Prompt from Standard Input
func GetPrompt() (*Prompt, error) {
	reader := bufio.NewReader(os.Stdin)
	promptStr, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	promptStr = strings.TrimSpace(promptStr)

	argSplit := strings.SplitN(promptStr, " ", 2)
	command := argSplit[0]
	arg := ""
	if len(argSplit) > 1{
		arg = argSplit[1]
	}

	return &Prompt{
		Command: command,
		Arg: arg,
	}, nil
}
