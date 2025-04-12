package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT" // Identifiers and literals
	INT       = "INT"
	ASSIGN    = "=" // Operators
	PLUS      = "+"
	MINUS     = "-"
	DIVIDE    = "/"
	ASTERISK  = "*"
	BANG      = "!"
	COMMA     = "," // Delimiters
	SEMICOLON = ";"
	LPARAN    = "(" // Braces
	RPARAN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION" // Keywords
	LET       = "LET"
	LT        = "<"
	GT        = ">"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIndent(indent string) TokenType {
	if tok, ok := keywords[indent]; ok {
		return tok
	}
	return IDENT
}
