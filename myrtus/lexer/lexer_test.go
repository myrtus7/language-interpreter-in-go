package lexer

import (
	"fmt"
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
