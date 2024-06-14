package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Builtin struct {
	command string
	handler func(context Context)
}

type Context struct {
	args     []string
	builtins []Builtin
	paths    []string
	wd       string
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
	wd, _ := os.Getwd()

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
		{
			command: "pwd",
			handler: func(context Context) {
				fmt.Fprintf(os.Stdout, "%s\n", context.wd)
			},
		},
		{
			command: "cd",
			handler: func(context Context) {
				if len(context.args) != 2 {
					fmt.Fprintf(os.Stdout, "cd: wrong number of arguments\n")
					return
				}

				err := os.Chdir(context.args[1])
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", context.args[1])
				}
			},
		},
	}

	for _, h := range handlers {
		if h.command == command {
			h.handler(Context{
				args:     args,
				builtins: handlers,
				paths:    paths,
				wd:       wd,
			})
			return
		}
	}

	// not builtin, then try to find the command in the PATH
	for _, path := range paths {
		if _, err := os.Stat(path + "/" + command); err == nil {
			// command found
			// execute the program
			cmd := exec.Command(path+"/"+command, args[1:]...)
			cmd.Env = os.Environ()
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run() // Run starts the specified command and waits for it to complete.
			if err != nil {
				// handle error
				fmt.Println("Command finished with error: ", err)
			}
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
