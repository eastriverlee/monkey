package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	ID = "ID"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	FUNC = "FUNC"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"func": FUNC,
	"let": LET,
}

func LookupId(id string) TokenType {
	if token, ok := keywords[id]; ok {
		return token
	}
	return ID
}
