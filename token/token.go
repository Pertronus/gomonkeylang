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
	LT        = "<"
	GT        = ">"
	EQ        = "="
	NOT_EQ    = "!="
	COMMA     = "," // Delimiters
	SEMICOLON = ";"
	LPARAN    = "(" // Braces
	RPARAN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION" // Keywords
	LET       = "LET"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIndent(indent string) TokenType {
	if tok, ok := keywords[indent]; ok {
		return tok
	}
	return IDENT
}
