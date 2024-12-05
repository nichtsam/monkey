package lexer

import (
	"github.com/nichtsam/monkey/token"
)

type Lexer struct {
	input        string
	char         byte
	position     int
	readPosition int
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) readChar() {
	l.char = l.peekChar()
	if l.char == '0' {
		return
	}

	l.forwardPosition()
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) forwardPosition() {
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readLetters() string {
	start := l.position
	for isLetter(l.peekChar()) {
		l.forwardPosition()
	}
	// l.char is outdated now
	return l.input[start : l.position+1]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.peekChar()) {
		l.forwardPosition()
	}
	// l.char is outdated now
	return l.input[start : l.position+1]
}

func (l *Lexer) skipWhitespaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.readChar()
	l.skipWhitespaces()

	switch l.char {
	case ',':
		t = token.New(token.COMMA, string(l.char))
	case ';':
		t = token.New(token.SEMICOLON, string(l.char))
	case '(':
		t = token.New(token.L_PAREN, string(l.char))
	case ')':
		t = token.New(token.R_PAREN, string(l.char))
	case '{':
		t = token.New(token.L_BRACE, string(l.char))
	case '}':
		t = token.New(token.R_BRACE, string(l.char))
	case '=':
		if l.peekChar() == '=' {
			start := l.position
			l.forwardPosition()
			t = token.New(token.EQUAL, l.input[start:l.position+1])
		} else {
			t = token.New(token.ASSIGN, string(l.char))
		}
	case '+':
		t = token.New(token.PLUS, string(l.char))
	case '-':
		t = token.New(token.MINUS, string(l.char))
	case '!':
		if l.peekChar() == '=' {
			start := l.position
			l.forwardPosition()
			t = token.New(token.NOT_EQUAL, l.input[start:l.position+1])
		} else {
			t = token.New(token.BANG, string(l.char))
		}
	case '*':
		t = token.New(token.ASTERISK, string(l.char))
	case '/':
		t = token.New(token.SLASH, string(l.char))
	case '<':
		t = token.New(token.LT, string(l.char))
	case '>':
		t = token.New(token.GT, string(l.char))
	case 0:
		t = token.New(token.EOF, string(l.char))
	default:
		if isLetter(l.char) {
			letters := l.readLetters()
			if tokenType, ok := token.Keywords[letters]; ok {
				t = token.New(tokenType, letters)
			} else {
				t = token.New(token.IDENTIFIER, letters)
			}
		} else if isDigit(l.char) {
			t = token.New(token.INTERGER, l.readNumber())
		} else {
			t = token.New(token.ILLEGAL, string(l.char))
		}
	}

	return t
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
