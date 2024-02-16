package lexer

type character byte

func (ch character) isLetter() bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (ch character) isDigit() bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func (ch character) toString() string {
	return string(ch)
}
