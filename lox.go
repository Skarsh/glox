package main

import (
	"fmt"
	"log"
	"os"
)

type Lox struct {
	hadError bool
}

func runFile(filePath string) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error opening file")
		os.Exit(69)
	}
	run(string(fileBytes[:]))
}

func runPrompt() {
	for {
		fmt.Print("> ")
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			fmt.Println("Error")
			break
		}
		run(line)
	}
}

func run(source string) {
	scanner := Scanner{source}
	scanner.scanTokens()
}

func error(line int, message string) {

}

func Report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}
