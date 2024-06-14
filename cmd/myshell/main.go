package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	args := strings.Split(input, " ")

	if args[0] == "exit\n" {
		os.Exit(0)
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found", args[0])
	}
}
