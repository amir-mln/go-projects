package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewBasicToken(str string) Token {
	if str == NULL_STRING {
		return Token{Type: EOF, Literal: str}
	}

	types := map[string]TokenType{
		"=":  ASSIGN,
		"==": EQUAL,
		"!=": UNEQUAL,
		"!":  EXCLAMATION,
		";":  SEMICOLON,
		",":  COMMA,
		"*":  ASTERISK,
		"/":  SLASH,
		"+":  PLUS,
		"-":  MINUS,
		"(":  LPAREN,
		")":  RPAREN,
		"{":  LBRACE,
		"}":  RBRACE,
		"<":  LCHEVRON,
		">":  RCHEVRON,
	}

	if t, ok := types[str]; ok {
		return Token{Type: t, Literal: str}
	}

	return Token{Type: ILLEGAL, Literal: str}
}

func NewIdentToken(str string) Token {
	types := map[string]TokenType{
		"function": FUNCTION,
		"let":      LET,
		"true":     TRUE,
		"false":    FALSE,
		"if":       IF,
		"else":     ELSE,
		"return":   RETURN,
	}

	if t, ok := types[str]; ok {
		return Token{Type: t, Literal: str}
	}

	return Token{Type: IDENT, Literal: str}
}

func NewNumericToken(str string) Token {
	if strings.ContainsRune(str, '.') {
		return Token{Type: DECIMAL, Literal: str}
	} else {
		return Token{Type: INT, Literal: str}
	}
}

const (
	NULL_CHAR   = byte(0)
	NULL_STRING = string(byte(0))
)

const (
	ILLEGAL     TokenType = "ILLEGAL"
	EOF         TokenType = "EOF"
	IDENT       TokenType = "IDENT"
	ASSIGN      TokenType = "="
	EQUAL       TokenType = "=="
	UNEQUAL     TokenType = "!="
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	EXCLAMATION TokenType = "!"
	ASTERISK    TokenType = "*"
	SLASH       TokenType = "/"
	COMMA       TokenType = ","
	SEMICOLON   TokenType = ";"
	LPAREN      TokenType = "("
	RPAREN      TokenType = ")"
	LBRACE      TokenType = "{"
	RBRACE      TokenType = "}"
	LCHEVRON    TokenType = "<"
	RCHEVRON    TokenType = ">"

	INT     TokenType = "INT"
	DECIMAL TokenType = "DECIMAL"

	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)
