package parser

type Token struct {
	Value string
}

func Lexer(source string) []Token {
	println("Lexing: " + source)

	return []Token{{Value: "ast"}, {Value: "best"}}
}
