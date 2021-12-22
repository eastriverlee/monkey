package lexer

import "monkey/token"

type Lexer struct {
	input		string
	index		int
	nextIndex	int
	char		byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.nextIndex >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.nextIndex]
	}
	lexer.index = lexer.nextIndex
	lexer.nextIndex++
}

func (lexer *Lexer) NextToken() token.Token {
	var token_ token.Token

	lexer.skipSpaces()

	switch lexer.char {
	case '=':
		token_ = newToken(token.ASSIGN, lexer.char)
	case ';':
		token_ = newToken(token.SEMICOLON, lexer.char)
	case '(':
		token_ = newToken(token.LPAREN, lexer.char)
	case ')':
		token_ = newToken(token.RPAREN, lexer.char)
	case '{':
		token_ = newToken(token.LBRACE, lexer.char)
	case '}':
		token_ = newToken(token.RBRACE, lexer.char)
	case '[':
		token_ = newToken(token.LBRACK, lexer.char)
	case ']':
		token_ = newToken(token.RBRACK, lexer.char)
	case ',':
		token_ = newToken(token.COMMA, lexer.char)
	case '+':
		token_ = newToken(token.PLUS, lexer.char)
	case 0:
		token_.Literal = ""
		token_.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			token_.Literal = lexer.readId()
			token_.Type = token.LookupId(token_.Literal)
			return token_
		} else if isDigit(lexer.char) {
			token_.Type = token.INT
			token_.Literal = lexer.readNumber()
			return token_
		} else {
			token_ = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return token_
}

func (lexer *Lexer) readId() string {
	index := lexer.index
	for isLetter(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[index:lexer.index]
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func newToken(tokenType token.TokenType, c byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(c)}
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func (lexer *Lexer) skipSpaces() {
	for isSpace(lexer.char) {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string {
	index := lexer.index

	println("number")
	for isDigit(lexer.char) {
		println(string(lexer.char))
		lexer.readChar()
	}
	return lexer.input[index:lexer.index]
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
