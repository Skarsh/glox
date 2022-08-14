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
	if s.current+1 < utf8.RuneCountInString(s.source) {
		return rune(s.source[s.current]), nil
	}
	return ' ', errors.New("error, trying to advance beyond the length of source")
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal *Object) {
	text := s.source[s.start:s.current]
	s.tokens.PushBack(NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) scanToken() {
	c, err := s.advance()
	if err != nil {
		// TODO: Should this be logged instead?
		fmt.Println("Error in scanning token")
	}
	switch c {
	case '(':
		s.addToken(LeftParen)
		break
	case ')':
		s.addToken(RightParen)
		break
	case '{':
		s.addToken(LeftBrace)
		break
	case '}':
		s.addToken(RightBrace)
		break
	case ',':
		s.addToken(Comma)
		break
	case '.':
		s.addToken(Dot)
		break
	case '-':
		s.addToken(Minus)
		break
	case '+':
		s.addToken(Plus)
		break
	case ';':
		s.addToken(Semicolon)
		break
	case '*':
		s.addToken(Star)
		break
	default:
		LoxError(s.line, "Unexpected character.")
		break
	}
}

func (s *Scanner) scanTokens() list.List {
	tokens := list.New()
	fmt.Println("source:\n", s.source)

	return *tokens
}
