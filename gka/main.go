package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"gka.com/front-end"
	"gka.com/tools"
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
	tokens := frontend.ScanSource([]rune(source))
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
	var generateGrammar bool
	flag.BoolVar(&generateGrammar, "gen-grammar", false, "Generate grammar")
	flag.Parse()

	args := flag.Args()

	if generateGrammar {
		tools.GenerateGrammarFiles()
		return
	}

	args = os.Args
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
