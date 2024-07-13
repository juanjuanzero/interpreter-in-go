package token

const (
	ILLEGAL = "ILLEGAL" // characters of what we really dont know
	EOF     = "EOF"     // end of file

	// identifiers
	IDENT = "IDENT" // function names, variables and things like that
	INT   = "INT"   // int literal

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
