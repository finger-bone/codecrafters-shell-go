package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func oneShot() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	command := args[0]

	if command == "exit" {
		c, _ := strconv.Atoi(args[1])
		os.Exit(
			c,
		)
	} else if command == "echo" {
		fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args[1:], " "))
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", args[0])
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		oneShot()
	}
}
