package lexer

/*
* A lexer scans and evaluates source code, recognizes reserved identifiers,
* discarding whitespace, and turning meaningless strings into lists of tokens.
* It is a validator for source code. The lexer ingests streams of text from the
* source code files, converts them into usable tokens and sends them to a parser
 */

import (
	"github.com/IAmRiteshKoushik/hermes/token"
)

type Lexer struct {
	input        string
	position     int // current char pointer
	readPosition int // next char pointer (useful in PEEK)
	ch           byte
}

func New(input string) *Lexer {
	/*
	 * Initialization is happening by reading of the first character, otherwise the
	 * struct receives a {"input-string" 0 0 0} value which runs into EOF the moment
	 * we try to read the nextToken as the first token has not been read
	 * */
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// TODO: Try supporting the entire Unicode (and emojis) in the lexer
func (l *Lexer) readChar() {
	// If we have reached the end of the input, then
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII char for NULL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
