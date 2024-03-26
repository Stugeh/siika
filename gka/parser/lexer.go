package parser

import "gka.com/logger"

func charAtIndexEq(idx int, source string, cmp string) bool {
	return string(source[idx]) == cmp
}

func scanToken(start *int, current *int, lineNumber int, source string, tokens *[]Token) {
	lexeme := string(source[*start])

	switch lexeme {
	// *** SINGLE CHAR LEXEMES
	case "(":
		*tokens = append(*tokens, CreateToken(LEFT_PAREN, lexeme, "", lineNumber))
	case ")":
		*tokens = append(*tokens, CreateToken(RIGHT_PAREN, lexeme, "", lineNumber))
	case "{":
		*tokens = append(*tokens, CreateToken(LEFT_BRACE, lexeme, "", lineNumber))
	case "}":
		*tokens = append(*tokens, CreateToken(RIGHT_BRACE, lexeme, "", lineNumber))
	case ",":
		*tokens = append(*tokens, CreateToken(COMMA, lexeme, "", lineNumber))
	case ".":
		*tokens = append(*tokens, CreateToken(DOT, lexeme, "", lineNumber))

	case "-":
		*tokens = append(*tokens, CreateToken(MINUS, lexeme, "", lineNumber))
	case "+":
		*tokens = append(*tokens, CreateToken(PLUS, lexeme, "", lineNumber))
	case ";":
		*tokens = append(*tokens, CreateToken(SEMICOLON, lexeme, "", lineNumber))
	case "*":
		*tokens = append(*tokens, CreateToken(STAR, lexeme, "", lineNumber))

	// Lexemes that potentially contain more than one char
	case "!":
		if charAtIndexEq(*current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(BANG_EQUAL, "!=", "", lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(BANG, "!", "", lineNumber))
		}
	case "=":
		if charAtIndexEq(*current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(EQUAL_EQUAL, "==", "", lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(EQUAL, "=", "", lineNumber))
		}
	case "<":
		if charAtIndexEq(*current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(LESS_EQUAL, "<=", "", lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(LESS, "<", "", lineNumber))

		}

	case ">":
		if charAtIndexEq(*current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(GREATER_EQUAL, ">=", "", lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(GREATER, ">", "", lineNumber))

		}

	// Slash
	case "/":
		if charAtIndexEq(*current+1, source, "/") {
			// Ignore comments until end of line
			for *current < len(source) && source[*current] != '\n' {
				*current++
			}
		} else {
			*tokens = append(*tokens, CreateToken(SLASH, "/", "", lineNumber))
		}

	// ignore whitespace
	case " ", "\r", "\t":

	default:
		logger.Error(lineNumber, "Unexpected character."+lexeme)
	}

}

func ScanSource(source string) []Token {
	tokens := []Token{}
	start, line := 0, 1

	for current := 0; current < len(source); current++ {
		// Set start of token to current index
		start = current
		scanToken(&start, &current, line, source, &tokens)
	}

	return tokens
}
