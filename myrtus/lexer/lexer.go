package lexer

import (
	"myrtus/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar() //position to first char
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) generateToken() token.Token {
	var tok token.Token

	switch l.char {
	case '{':
		tok = createToken(token.LEFT_BRACE, l.char)
	case '}':
		tok = createToken(token.RIGHT_BRACE, l.char)
	case '(':
		tok = createToken(token.LEFT_PAREN, l.char)
	case ')':
		tok = createToken(token.RIGHT_PAREN, l.char)
	case '=':
		tok = createToken(token.ASSIGN, l.char)
	case '+':
		tok = createToken(token.PLUS, l.char)
	case ';':
		tok = createToken(token.SEMICOLON, l.char)
	case ',':
		tok = createToken(token.COMMA, l.char)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		tok = createToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return tok
}

func createToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
