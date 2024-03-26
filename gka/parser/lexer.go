package parser

import "gka.com/logger"

func charAtIndexEq(idx int, source []rune, cmp string) bool {
	return string(source[idx]) == cmp
}

func scanToken(start *int, current *int, lineNumber *int, source []rune, tokens *[]Token) {
	lexeme := []rune{source[*start]}

	switch lexeme[0] {
	case '(', ')', '{', '}', ',', '.', '-', '+', ';', '*':
		*tokens = append(*tokens, CreateToken(getTokenType(lexeme[0]), lexeme, "", *lineNumber))

	case '!', '=', '<', '>':
		if charAtIndexEq(*current+1, source, "=") {
			lexeme = append(lexeme, '=')
			*tokens = append(*tokens, CreateToken(getTokenType(lexeme[0]), lexeme, "", *lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(getTokenType(lexeme[0]), lexeme, "", *lineNumber))
		}

	case '/':
		if charAtIndexEq(*current+1, source, "/") {
			for *current < len(source) && source[*current] != '\n' {
				*current++
			}
			*lineNumber++
		} else {
			*tokens = append(*tokens, CreateToken(SLASH, lexeme, "", *lineNumber))
		}

	case ' ', '\r', '\t':

	case '\n':
		*lineNumber++

	default:
		logger.Error(*lineNumber, "Unexpected character."+string(lexeme))
	}
}

func getTokenType(char rune) TokenType {
	switch char {
	case '(':
		return LEFT_PAREN
	case ')':
		return RIGHT_PAREN
	case '{':
		return LEFT_BRACE
	case '}':
		return RIGHT_BRACE
	case ',':
		return COMMA
	case '.':
		return DOT
	case '-':
		return MINUS
	case '+':
		return PLUS
	case ';':
		return SEMICOLON
	case '*':
		return STAR
	case '!':
		return BANG
	case '=':
		return EQUAL
	case '<':
		return LESS
	case '>':
		return GREATER
	default:
		return UNKNOWN
	}
}

func ScanSource(source []rune) []Token {
	tokens := []Token{}
	start, line := 0, 1

	for current := 0; current < len(source); current++ {
		// Set start of token to current index
		start = current
		scanToken(&start, &current, &line, source, &tokens)
	}

	return tokens
}
