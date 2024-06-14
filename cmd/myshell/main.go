package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Handler struct {
	command string
	handler func(context Context)
}

type Context struct {
	args     []string
	handlers []Handler
}

func oneShot() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	command := args[0]

	handlers := []Handler{
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

	found := false

	for _, h := range handlers {
		if h.command == command {
			h.handler(Context{
				args:     args,
				handlers: handlers,
			})
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		oneShot()
	}
}
