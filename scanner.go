package main

import (
	"container/list"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Scanner struct {
	source  string
	tokens  *list.List
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{source: source, tokens: list.New(), start: 0, current: 0, line: 1}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= utf8.RuneCountInString(s.source)
}

func (s *Scanner) advance() (rune, error) {
	s.current += 1
	if s.current + 1 < utf8.RuneCountInString(s.source) {
		return rune(s.source[s.current]), nil
	}
	return ' ', errors.New("error, trying to advance beyond the length of source")
}

func (s *Scanner) addToken(tokenType TokenType) {

}

func (s *Scanner) scanTokens() list.List {
	tokens := list.New()
	fmt.Println("source:\n", s.source)
	return *tokens
}
