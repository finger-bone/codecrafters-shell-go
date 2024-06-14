package main

import (
	"fmt"
	"os"
	"strings"
)

func cd(context Context) {
	if len(context.args) != 2 {
		fmt.Fprintf(os.Stdout, "cd: wrong number of arguments\n")
		return
	}

	to := strings.ReplaceAll(
		context.args[1],
		"~",
		os.Getenv("HOME"),
	)

	err := os.Chdir(to)
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", context.args[1])
	}
}
