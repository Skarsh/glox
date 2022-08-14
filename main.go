package main

import (
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if argLength == 1 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
