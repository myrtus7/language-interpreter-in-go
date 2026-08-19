package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// enum like const for token types
const (
	FUNCTION = "FUNCTION"
	IDENT    = "IDENT"
	LET      = "LET"
	INT      = "INT"

	ASSIGN      = "="
	PLUS        = "+"
	RIGHT_PAREN = ")"
	LEFT_PAREN  = "("
	RIGHT_BRACE = "}"
	LEFT_BRACE  = "{"
	COMMI       = ","
	SEMICOLON   = ";"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)
