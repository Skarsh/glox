package main

import (
	"testing"
)

func TestReport(t *testing.T) {
	Report(2, "where", "message")
}

func TestNewScanner(t *testing.T) {
	expectedSource := "hallo"
	scanner := NewScanner(expectedSource)
	if scanner.source != expectedSource {
		t.Errorf("source = %s, want %s", scanner.source, expectedSource)
	}

	tokensLen := scanner.tokens.Len()
	expectedTokensLen := 0
	if tokensLen != expectedTokensLen {
		t.Errorf("tokens len = %d, want %d", tokensLen, expectedTokensLen)
	}
}

func TestIsAtEnd(t *testing.T) {

	scanner := NewScanner("test")
	if scanner.isAtEnd() == true {
		t.Errorf("scanner should not be at the end of source")
	}

	scanner.source = ""
	if scanner.isAtEnd() == false {
		t.Errorf("scanner should be at the end of source")
	}
}

func TestAdvance(t *testing.T) {
	scanner := NewScanner("test")
	expectedRune := 'e'
	nextRune, err := scanner.advance()
	if err != nil {
		t.Errorf("scanner.advance() returned error %s", err)
	}
	if nextRune != expectedRune {
		t.Errorf("advance() produced = %c, want %c", nextRune, expectedRune)
	}
}
