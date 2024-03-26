package parser

func scanToken(start int, current int, source string, tokens *[]Token) {
	lexeme := string(source[start:current])

	switch lexeme {
	// *** SINGLE CHAR LEXEMES
	case "(":
		*tokens = append(*tokens, CreateToken(LEFT_PAREN, lexeme))
		break
	case ")":
		*tokens = append(*tokens, CreateToken(RIGHT_PAREN, lexeme))
		break
	case "{":
		*tokens = append(*tokens, CreateToken(LEFT_BRACE, lexeme))
		break
	case "}":
		*tokens = append(*tokens, CreateToken(RIGHT_BRACE, lexeme))
		break
	case ",":
		*tokens = append(*tokens, CreateToken(COMMA, lexeme))
		break
	case ".":
		*tokens = append(*tokens, CreateToken(DOT, lexeme))
		break

	case "-":
		*tokens = append(*tokens, CreateToken(MINUS, lexeme))
		break
	case "+":
		*tokens = append(*tokens, CreateToken(PLUS, lexeme))
		break
	case ";":
		*tokens = append(*tokens, CreateToken(SEMICOLON, lexeme))
		break
	case "*":
		*tokens = append(*tokens, CreateToken(STAR, lexeme))
		break
	}
}

func ScanSource(source string) []Token {
	endReached := false
	tokens := []Token{}
	start, current, line := 0, 0, 1

	// While we havent reached the end of source
	for !(current >= len(source)) {
		start = current

		scanToken(start, current, source, &tokens)

	}

	return []Token{{Lexeme: "lexeme", Literal: "literal", Line: 10}}
}
