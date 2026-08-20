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
	COMMA     = ","
	SEMICOLON   = ";"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"let" : LET,
}
func GetKeywordOrIdentType(literal string) TokenType{
	if keyword, ok := keywords[literal]; ok {
		return keyword
	}
	return IDENT

}