package token

type (
	TokenType string
	Token     struct {
		Type    TokenType
		Literal string
	}
)

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func New(tokenType TokenType, literal string) Token {
	return Token{tokenType, literal}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literal
	IDENTIFIER = "IDENTIFIER"
	INTERGER   = "INTERGER"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	L_PAREN   = "("
	R_PAREN   = ")"
	L_BRACE   = "{"
	R_BRACE   = "}"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)
