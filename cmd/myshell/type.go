package main

import (
	"fmt"
	"os"
)

func Type(context Context) {
	command := context.args[1]

	for _, h := range context.handlers {
		if h.command == command {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
			return
		}
	}

	fmt.Fprintf(os.Stdout, "%s: not found\n", command)
}
