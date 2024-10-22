package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Token or character that we don't know
	EOF     = "EOF"     // End of file

	// Idenfitifers and literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	/* As identifiers can be keywords and non-keywords, for distinguishing them
	 * this function will lookup and if it can find the keyword then the
	 * corresponding constant is returned otherwise token.IDENT is the default
	 */
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
