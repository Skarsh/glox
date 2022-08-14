package main

import (
	"testing"
)

func TestNumTokens(t *testing.T) {
	var eof TokenType
	eof = Eof

	// The idea here is that the Eof token is always the last Token
	// in the TokenType enumeration and that the enum starts at value 1
	// e.g. iota + 1
	if eof != NumTokens {
		t.Errorf("eof = %d, want %d", eof, NumTokens)
	}
}

func TestTokenToString(t *testing.T) {
	var tokenType TokenType
	tokenType = Star

	lexeme := "lexeme"

	literal := Object{data: "data123"}
	line := 1

	token := Token{
		tokenType,
		lexeme,
		literal,
		line,
	}

	tokenString := token.String()
	expectedTokenString := "Star lexeme {data123}"
	if tokenString != expectedTokenString {
		t.Errorf("token.String = %s, want Star lexeme %s", tokenString, expectedTokenString)
	}
}
