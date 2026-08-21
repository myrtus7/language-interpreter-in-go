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

	l.eatWhiteSpaces()

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
		if l.peekChar() == '=' {
			tok = l.createTwoCharToken(l.peekChar(), token.EQ)
		} else {
			tok = createToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = createToken(token.PLUS, l.char)
	case ';':
		tok = createToken(token.SEMICOLON, l.char)
	case ',':
		tok = createToken(token.COMMA, l.char)
	case '-':
		if isDigit(l.peekChar()) {
			l.readChar()
			literal := "-" + l.readNumber()
			return token.Token{Type: token.INT, Literal: literal}
		} else {
			tok = createToken(token.MINUS, l.char)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.createTwoCharToken(l.peekChar(), token.NOT_EQ)
		} else {
			tok = createToken(token.BANG, l.char)
		}

	case '*':
		tok = createToken(token.ASTERISK, l.char)
	case '/':
		tok = createToken(token.SLASH, l.char)
	case '<':
		tok = createToken(token.LT, l.char)
	case '>':
		tok = createToken(token.GT, l.char)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.GetKeywordOrIdentType(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = createToken(token.ILLEGAL, l.char)
		}

		return tok
	}

	l.readChar()
	return tok
}

func createToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) eatWhiteSpaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) createTwoCharToken(peekedChar byte, tokenType token.TokenType) token.Token {
	literal := string(l.char) + string(peekedChar)
	l.readChar()
	return token.Token{Type: tokenType, Literal: literal}
}
