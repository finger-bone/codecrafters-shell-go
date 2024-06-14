package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

type Builtin struct {
	command string
	handler func(context Context)
}

type Context struct {
	args     []string
	builtins []Builtin
	paths    []string
}

func oneShot() {
	fmt.Fprint(os.Stdout, "$ ")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	command := args[0]
	paths := strings.Split(os.Getenv("PATH"), ":")

	handlers := []Builtin{
		{
			command: "exit",
			handler: Exit,
		},
		{
			command: "echo",
			handler: Echo,
		},
		{
			command: "type",
			handler: Type,
		},
	}

	for _, h := range handlers {
		if h.command == command {
			h.handler(Context{
				args:     args,
				builtins: handlers,
				paths:    paths,
			})
			return
		}
	}

	// not builtin, then try to find the command in the PATH
	for _, path := range paths {
		if _, err := os.Stat(path + "/" + command); err == nil {
			// command found
			// execute the program
			// use syscall.Exec
			syscall.Exec(path+"/"+command, args, os.Environ())
			return
		}
	}

	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		oneShot()
	}
}
