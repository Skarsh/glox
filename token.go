package main

import (
	"fmt"
)

type Object struct {
	data any
}

type TokenType int

func (t TokenType) String() string {
	return fmt.Sprintf("")
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   Object
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal Object, line int) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t *Token) String() string {
	// TODO: Do proper string formatting
	// return fmt.Sprintf("%b", t)
	return ""
}
