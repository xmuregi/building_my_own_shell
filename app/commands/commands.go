package commands


// Allowed Shell builtins
var commandList map[string]bool = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

// Returns bool value for Shell builtins
func IsCommand(command string) (bool) {
	return commandList[command]
}
