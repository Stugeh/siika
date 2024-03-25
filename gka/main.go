package main

import (
	"bufio"
	"fmt"
	"os"

	"gka.com/parser"
)

func readFile(filepath string) string {
	println("Running file: " + filepath)
	content, err := os.ReadFile(filepath)

	if err != nil {
		println(err.Error())
		return ""
	}

	return string(content[:])
}

func run(source string) {
	tokens := parser.Lexer(source)
	for _, token := range tokens {
		println(token.Value)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil || line == "" {
			continue
		}

		content := readFile(line)

		if len(content) > 0 {
			break
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
		content := readFile(args[0])

		if len(content) == 0 {
			println("failed to read input file")
			return
		}

		run(content)
	}

	runPrompt()
}
