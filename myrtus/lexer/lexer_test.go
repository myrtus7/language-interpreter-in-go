package lexer

import (
	"fmt"
	"testing"
	"myrtus/token"
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
