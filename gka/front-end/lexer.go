package frontend

import (
	"gka.com/logger"
)

func charAtIndexEq(idx int, source []rune, cmp string) bool {
	return string(source[idx]) == cmp
}

func scanToken(start *int, current *int, lineNumber *int, source []rune, tokens *[]Token) {
	lexeme := []rune{source[*start]}

	switch lexeme[0] {
	// *** SINGLE CHAR LEXEMES
	case '(':
		*tokens = append(*tokens, CreateToken(LEFT_PAREN, lexeme, "", *lineNumber))
	case ')':
		*tokens = append(*tokens, CreateToken(RIGHT_PAREN, lexeme, "", *lineNumber))
	case '{':
		*tokens = append(*tokens, CreateToken(LEFT_BRACE, lexeme, "", *lineNumber))
	case '}':
		*tokens = append(*tokens, CreateToken(RIGHT_BRACE, lexeme, "", *lineNumber))
	case ',':
		*tokens = append(*tokens, CreateToken(COMMA, lexeme, "", *lineNumber))
	case '.':
		*tokens = append(*tokens, CreateToken(DOT, lexeme, "", *lineNumber))

	case '-':
		*tokens = append(*tokens, CreateToken(MINUS, lexeme, "", *lineNumber))
	case '+':
		*tokens = append(*tokens, CreateToken(PLUS, lexeme, "", *lineNumber))
	case ';':
		*tokens = append(*tokens, CreateToken(SEMICOLON, lexeme, "", *lineNumber))
	case '*':
		*tokens = append(*tokens, CreateToken(STAR, lexeme, "", *lineNumber))

	// Lexemes that potentially contain more than one char
	case '!':
		if charAtIndexEq(*current+1, source, "=") {
			lexeme = append(lexeme, '=')
			*tokens = append(*tokens, CreateToken(BANG_EQUAL, lexeme, "", *lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(BANG, lexeme, "", *lineNumber))
		}
	case '=':
		if charAtIndexEq(*current+1, source, "=") {
			lexeme = append(lexeme, '=')
			*tokens = append(*tokens, CreateToken(EQUAL_EQUAL, lexeme, "", *lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(EQUAL, lexeme, "", *lineNumber))
		}
	case '<':
		if charAtIndexEq(*current+1, source, "=") {
			lexeme = append(lexeme, '=')
			*tokens = append(*tokens, CreateToken(LESS_EQUAL, lexeme, "", *lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(LESS, lexeme, "", *lineNumber))

		}

	case '>':
		if charAtIndexEq(*current+1, source, "=") {
			lexeme = append(lexeme, '=')
			*tokens = append(*tokens, CreateToken(GREATER_EQUAL, lexeme, "", *lineNumber))
			*current++
		} else {
			*tokens = append(*tokens, CreateToken(GREATER, lexeme, "", *lineNumber))

		}

	// Slash
	case '/':
		if charAtIndexEq(*current+1, source, "/") {
			// Ignore comments until end of line
			for *current < len(source) && source[*current] != '\n' {
				*current++
			}
			*lineNumber++
		} else {
			*tokens = append(*tokens, CreateToken(SLASH, lexeme, "", *lineNumber))
		}

	// ignore whitespace
	case ' ', '\r', '\t':

	case '\n':
		*lineNumber++

	// Strings
	case '"':
		lexeme = lexeme[1:] // Slice off opening quote

		*current++
		for *current < len(source) && source[*current] != '\n' && source[*current] != '"' {
			lexeme = append(lexeme, source[*current])
			*current++
		}

		if *current >= len(source) || source[*current] != '"' {
			logger.Error(*lineNumber, "Unterminated string literal.")
			break
		}

		*tokens = append(*tokens, CreateToken(STRING, lexeme, "", *lineNumber))

	// Key words
	case 'o':
		if source[*current+1] == 'r' {
			*current++
			*tokens = append(*tokens, CreateToken(OR, source[*start:*current], "", *lineNumber))
		}

	default:
		// handle numbers
		// TODO make logic less shit
		if isDigit(source[*current]) {
			*current++
			for *current < len(source) && isDigit(source[*current]) {
				*current++
			}

			// handle floating point nums
			if *current < len(source) && source[*current] == '.' && isDigit(source[*current+1]) {
				for *current < len(source) && isDigit(source[*current+1]) {
					*current++
				}
				*tokens = append(*tokens, CreateToken(NUMBER, source[*start:*current+1], "", *lineNumber))
			} else {
				*current--
				*tokens = append(*tokens, CreateToken(NUMBER, source[*start:*current+1], "", *lineNumber))
			}

			return

		} else if isAlpha(source[*current]) { // Handle key words
			for isAlphaNumeric(source[*current+1]) {
				*current++
			}

			tokenType := MatchKeyword(source[*start:*current+1], *lineNumber)
			if tokenType == UNKNOWN {
				return
			}
			*tokens = append(*tokens, CreateToken(tokenType, source[*start:*current+1], "", *lineNumber))

			return
		}

		logger.Error(*lineNumber, "Unexpected character."+string(lexeme))
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

// ** Helper methods **//
// Simpler than unicode.IsDigit()
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isAlphaNumeric(char rune) bool {
	return isAlpha(char) || isDigit(char)
}
