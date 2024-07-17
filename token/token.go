package token

const (
	ILLEGAL = "ILLEGAL" // characters of what we really dont know
	EOF     = "EOF"     // end of file

	// identifiers
	IDENT = "IDENT" // function names, variables and things like that
	INT   = "INT"   // int literal

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	ASTERISK = "*"
	MINUS    = "-"
	SLASH    = "/"
	BANG     = "!"

	// Comparisons
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
