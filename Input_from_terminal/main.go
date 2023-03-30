package main

import (
	"os"
	"strings"
)

func greeter(name string) string {
	return "Hello " + strings.TrimSpace(name) + ", good morning"
}

func main() {
	args := os.Args

	print()

	if len(args[1]) != 2 {
		println("Invalid entries, please insert only your name")
		return
	}
	println(greeter(args[1]))
}