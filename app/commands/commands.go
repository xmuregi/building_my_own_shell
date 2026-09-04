package commands



var commandList map[string]bool = map[string]bool{
	"exit": true,
	"echo": true,
}

func IsCommand(command string) (bool, error) {
	return commandList[command], nil
}
