package lexer

import (
	"github.com/amir-mln/go-projects/pishi/token"
)

type Lexer struct {
	input        string
	position     int       // points to current character
	readPosition int       // points to the next character that we want to read
	ch           character // the character the 'position' point to
}

func (l *Lexer) peekChar() character {
	if l.readPosition >= len(l.input) {
		return character(token.NULL_CHAR)
	}

	return character(l.input[l.readPosition])
}

func (l *Lexer) readChar() character {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1

	return l.ch
}

func (l *Lexer) readBack() character {
	if l.readPosition != 0 {
		l.readPosition -= 1
	}
	if l.position != 0 {
		l.position -= 1
	}

	l.ch = character(l.input[l.position])

	return l.ch
}

func (l *Lexer) readIdentifier() string {
	ident := ""

	for l.ch.isLetter() {
		ident += l.ch.toString()
		l.readChar()
	}

	if len(ident) > 0 {
		l.readBack()
	}

	return ident
}

func (l *Lexer) readNumber() string {
	num := ""

	for l.ch.isDigit() {
		num += l.ch.toString()
		l.readChar()
	}

	if len(num) > 0 {
		l.readBack()
	}

	if num == "." {
		num = "0.0"
	}

	if num[0] == '.' {
		num = "0" + num
	}

	if num[len(num)-1] == '.' {
		num = num + "0"
	}

	return num
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpaces()

	switch {
	case l.ch.isLetter():
		tok = token.NewIdentToken(l.readIdentifier())
	case l.ch.isDigit():
		tok = token.NewNumericToken(l.readNumber())
	default:
		literal := l.ch.toString()

		if l.peekChar() == '=' && (l.ch == '=' || l.ch == '!') {
			l.readChar()
			literal += l.ch.toString()
		}

		tok = token.NewBasicToken(literal)
	}

	l.readChar()

	return tok
}

func New(input string) *Lexer {
	l := &Lexer{input: input, ch: character(input[0])}

	return l
}
