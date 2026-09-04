package commands

import (
	"fmt"
	"os"
)

func Echo(arg string) {
	fmt.Fprintf(os.Stdout, "%s\n", arg)
}
