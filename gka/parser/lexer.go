package parser

import "gka.com/logger"

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
		*current++
		for *current < len(source) && source[*current] != '\n' && source[*current] != '"' {
			lexeme = append(lexeme, source[*current])
			*current++
		}
		if *current >= len(source) || source[*current] != '"' {
			logger.Error(*lineNumber, "Unterminated string literal.")
			break
		}
		lexeme = append(lexeme, source[*current]) // Include the closing quote
		*tokens = append(*tokens, CreateToken(STRING, lexeme, "", *lineNumber))

	default:

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
