package main

import (
	"fmt"
)

type Scanner struct {
	source string
}

func (s *Scanner) scanTokens() {
	fmt.Println("source:\n", s.source)
}
