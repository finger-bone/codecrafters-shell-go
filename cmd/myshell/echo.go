package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo(context Context) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(context.args[1:], " "))
}
