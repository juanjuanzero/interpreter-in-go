package lexer

import "github.com/juanjuanzero/monkey/token"

type Lexer struct {
	input        string
	position     int // the current position in input, points to the current char
	readPosition int // the current reading position in input after the current char
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// initialize at the first char
	l.readChar()
	return l
}

// helper methods
func (l *Lexer) readChar() {
	// if the read position is greater than or equal to the byte we are out of bounds on the input
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCI code for NUL
	} else {
		// otherwise set the character under examination at the readPosition
		l.ch = l.input[l.readPosition]
	}
	// move the position forward
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// if the char is a letter, read for an identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar() // advance to the next char
	return tok
}

// checks if this is an ascii letter
func isLetter(ch byte) bool {
	// true if the code point (rune) is between a to z or A to Z or equal to _ for snake case
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	// save the current position
	position := l.position
	// continue reading until you reach a non char
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
