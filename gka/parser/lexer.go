package parser

import "gka.com/logger"

func charAtIndexEq(idx int, source string, cmp string) bool {
	return string(source[idx]) == cmp
}

func scanToken(start int, current int, lineNumber int, source string, tokens *[]Token) {
	lexeme := string(source[current])

	switch lexeme {
	// *** SINGLE CHAR LEXEMES
	case "(":
		*tokens = append(*tokens, CreateToken(LEFT_PAREN, lexeme, "", lineNumber))
		break
	case ")":
		*tokens = append(*tokens, CreateToken(RIGHT_PAREN, lexeme, "", lineNumber))
		break
	case "{":
		*tokens = append(*tokens, CreateToken(LEFT_BRACE, lexeme, "", lineNumber))
		break
	case "}":
		*tokens = append(*tokens, CreateToken(RIGHT_BRACE, lexeme, "", lineNumber))
		break
	case ",":
		*tokens = append(*tokens, CreateToken(COMMA, lexeme, "", lineNumber))
		break
	case ".":
		*tokens = append(*tokens, CreateToken(DOT, lexeme, "", lineNumber))
		break

	case "-":
		*tokens = append(*tokens, CreateToken(MINUS, lexeme, "", lineNumber))
		break
	case "+":
		*tokens = append(*tokens, CreateToken(PLUS, lexeme, "", lineNumber))
		break
	case ";":
		*tokens = append(*tokens, CreateToken(SEMICOLON, lexeme, "", lineNumber))
		break
	case "*":
		*tokens = append(*tokens, CreateToken(STAR, lexeme, "", lineNumber))
		break

	// Lexemes that potentially contain more than one char
	case "!":
		if charAtIndexEq(current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(BANG_EQUAL, "!=", "", lineNumber))
		} else {
			*tokens = append(*tokens, CreateToken(BANG, "!", "", lineNumber))
		}
		break
	case "=":
		if charAtIndexEq(current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(EQUAL_EQUAL, "==", "", lineNumber))
		} else {
			*tokens = append(*tokens, CreateToken(EQUAL, "=", "", lineNumber))
		}
		break
	case "<":
		if charAtIndexEq(current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(LESS_EQUAL, "<=", "", lineNumber))
		} else {
			*tokens = append(*tokens, CreateToken(LESS, "<", "", lineNumber))

		}
		break

	case ">":
		if charAtIndexEq(current+1, source, "=") {
			*tokens = append(*tokens, CreateToken(GREATER_EQUAL, ">=", "", lineNumber))
		} else {
			*tokens = append(*tokens, CreateToken(GREATER, ">", "", lineNumber))

		}
		break

	// Slash
	case "/":

		if charAtIndexEq(current+1, source, "=") {
		}
	default:
		logger.Error(lineNumber, "Unexpected character.")
	}
}

func ScanSource(source string) []Token {
	endReached := false
	tokens := []Token{}
	start, current, line := 0, 0, 1

	// While we havent reached the end of source
	for !(current >= len(source)) {
		start = current

		scanToken(start, current, 0, source, &tokens)

	}

	return []Token{{Lexeme: "lexeme", Literal: "literal", Line: 10}}
}
