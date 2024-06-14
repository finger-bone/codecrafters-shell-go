package main

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory(context Context) {
	fmt.Fprintf(os.Stdout, "%s\n", context.wd)
}
