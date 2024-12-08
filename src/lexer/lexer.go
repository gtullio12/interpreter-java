package lexer

import (
	"java/tokens"
)

type Lexer struct {
	value        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		value: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.value) {
		l.ch = 0
	} else {
		l.ch = l.value[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() (tok tokens.Token) {
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = tokens.Token{Type: tokens.PLUS, Literal: "+"}
	case '-':
		tok = tokens.Token{Type: tokens.MINUS, Literal: "-"}
	case '"':
		str := l.readString()
		tok = tokens.Token{Type: tokens.STRING, Literal: str}
	case '=':
		if l.peekChar() == '=' { // Was an `==`
			tok = tokens.Token{Type: tokens.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = tokens.Token{Type: tokens.ASSIGN, Literal: "="}
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = tokens.Token{Type: tokens.NOT_EQ, Literal: "!="}
		} else {
			tok = tokens.Token{Type: tokens.BANG, Literal: "!"}
		}
	case '.':
		tok = tokens.Token{Type: tokens.PERIOD, Literal: "."}
	case '*':
		tok = tokens.Token{Type: tokens.ASTERISK, Literal: "*"}
	case ',':
		tok = tokens.Token{Type: tokens.COMMA, Literal: ","}
	case ';':
		tok = tokens.Token{Type: tokens.SEMICOLON, Literal: ";"}
	case '[':
		tok = tokens.Token{Type: tokens.LSPAREN, Literal: "["}
	case ']':
		tok = tokens.Token{Type: tokens.RSPAREN, Literal: "]"}
	case '{':
		tok = tokens.Token{Type: tokens.LBRACE, Literal: "{"}
	case '}':
		tok = tokens.Token{Type: tokens.RBRACE, Literal: "}"}
	case '(':
		tok = tokens.Token{Type: tokens.LPAREN, Literal: "("}
	case ')':
		tok = tokens.Token{Type: tokens.RPAREN, Literal: ")"}
	case 0:
		tok = tokens.Token{Type: tokens.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := tokens.LookupIdentifier(literal)
			tok = tokens.Token{Type: tokenType, Literal: literal}
			return tok
		} else if isNumber(l.ch) {
			literal := l.readNumber()
			tok = tokens.Token{Type: tokens.INT, Literal: literal}
			return tok
		} else {
			tok = tokens.Token{Type: tokens.ILLEGAL, Literal: string(l.ch)}
			return tok
		}
	}

	l.readChar()
	return
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.value[position:l.position]

}

func (l *Lexer) readString() string {
	position := l.position
	l.readChar() // Move pointer forward so we are not looking at `"`
	for l.ch != '"' {
		l.readChar()
	}
	return l.value[position+1 : l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.value[position:l.position]
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') && c != ' '
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.value) {
		return 0
	} else {
		return l.value[l.readPosition]
	}
}
