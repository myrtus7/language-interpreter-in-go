package lexer

import (
	"fmt"
	"myrtus/token"
	"testing"
)

func TestReadingInputs(t *testing.T) {

	input := "gougou"

	lexer := New(input)

	for lexer.char != 0 {
		fmt.Printf("%c\n", lexer.char)
		lexer.readChar()
	}
}

func TestLexingCondensedInput(t *testing.T) {
	input := "+(}"

	lexer := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.LEFT_PAREN, "("},
		{token.RIGHT_BRACE, "}"},
	}

	for i, test := range tests {
		tok := lexer.generateToken()

		if tok.Type != test.expectedType {
			t.Fatalf("TYPE ERROR: at test %d. expected %q, got %q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("LITERAL ERROR: at test %d. expected %q, got %q", i, test.expectedLiteral, tok.Literal)
		}
	}

}

func TestLexingInput(t *testing.T) {
	input := `let five = 5; 
		let ten = 10;
        let add = fn(x, y) {
        x + y;
        };
        let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFT_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LEFT_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RIGHT_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		tok := l.generateToken()

		if tok.Type != test.expectedType {
			t.Fatalf("TYPE ERROR: at test %d. expected %q, got %q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("LITERAL ERROR: at test %d. expected %q, got %q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
