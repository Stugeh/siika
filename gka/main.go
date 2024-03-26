package main

import (
	"bufio"
	"fmt"
	"os"

	"gka.com/parser"
)

func runFile(filepath string) error {
	println("Running file: " + filepath)
	content, err := os.ReadFile(filepath)

	if err != nil {
		println(err.Error())
		println(filepath)
		panic("Failed to read file. Exiting.")
	}

	return run(string(content))
}

func run(source string) error {
	tokens := parser.ScanSource(source)
	for _, token := range tokens {
		println(token.String())
	}
	return nil
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if err != nil || line == "" {
			println("Couldn't read string or it was empty.")
			continue
		}

		error := run(line)

		if error != nil {
			println(error.Error())
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 2 {
		println("Usage: gka [script]")
		panic("Too many args")
	}

	if len(args) == 2 {
		runFile(args[1])

	} else {
		runPrompt()
	}

}
