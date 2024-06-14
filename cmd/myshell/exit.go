package main

import (
	"os"
	"strconv"
)

func Exit(context Context) {
	c, _ := strconv.Atoi(context.args[1])
	os.Exit(
		c,
	)
}
