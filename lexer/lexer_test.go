package lexer

import (
	"testing"

	"github.com/juanjuanzero/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lex := New(input)

	for i, textTest := range tests {
		tok := lex.NextToken() // TODO implement

		if tok.Type != textTest.expectedType {
			t.Fatalf("test[%d] - tokentype is wrong. expected =%q. got %q", i, textTest.expectedType, tok.Type)
		}
		if tok.Literal != textTest.expectedLiteral {
			t.Fatalf("test[%d] - literal is wrong. expected =%q. got %q", i, textTest.expectedLiteral, tok.Literal)
		}
	}
}
