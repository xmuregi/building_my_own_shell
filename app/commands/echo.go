package commands

import (
	"fmt"
	"os"
)

// Echoes back the argument given to it
func Echo(arg string) {
	fmt.Fprintf(os.Stdout, "%s\n", arg)
}
