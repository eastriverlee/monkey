package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(test *testing.T) {
	input := `
let five = 5;
let ten = 10;

let add = func(x, y) {
	x + y;
}
let result = add(five, ten);
	`

	validations := []struct{
		expectedType	token.TokenType
		expectedLiteral	string
	}{
		{token.LET,			"let"},
		{token.ID,			"five"},
		{token.ASSIGN,		"="},
		{token.INT,			"5"},
		{token.SEMICOLON,	";"},

		{token.LET,			"let"},
		{token.ID,			"ten"},
		{token.ASSIGN,		"="},
		{token.INT,			"10"},
		{token.SEMICOLON,	";"},

		{token.LET,			"let"},
		{token.ID,			"add"},
		{token.ASSIGN,		"="},
		{token.FUNC,		"func"},
		{token.LPAREN,		"("},
		{token.ID,			"x"},
		{token.COMMA,		","},
		{token.ID,			"y"},
		{token.RPAREN,		")"},

		{token.LBRACE,		"{"},
		{token.ID,			"x"},
		{token.PLUS,		"+"},
		{token.ID,			"y"},
		{token.SEMICOLON,	";"},
		{token.RBRACE,		"}"},

		{token.LET,			"let"},
		{token.ID,			"result"},
		{token.ASSIGN,		"="},
		{token.ID,			"add"},
		{token.LPAREN,		"("},
		{token.ID,			"x"},
		{token.COMMA,		","},
		{token.ID,			"y"},
		{token.RPAREN,		")"},
		{token.SEMICOLON,	";"},

		{token.EOF,			""},
	}

	lexer := NewLexer(input)

	for i, validation := range validations {
		token := lexer.NextToken()
		if token.Type != validation.expectedType {
			test.Fatalf("tests[%d] - wrong tokentype. expected=%q, got=%q", i, validation.expectedType, token.Type)
		}
	}
}
