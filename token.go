package main

import (
	"fmt"
)

type Object struct {
	data any
}

func (o *Object) String() string {
	return fmt.Sprintf("%s", o.data)
}

type TokenType int

const NumTokens = 39

// Enumeration describing all the different token types.
// Importatin invariant to uphold is that the first token
// should have enum value equal to 1, e.g. iota + 1,
// and the Eof token should always be the last. This gives
// and easy way of asserting that the number of possible tokens
// matches that which are specified in the NumTokens constant.
const (
	// Single-character tokens.
	LeftParen TokenType = iota + 1
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or two character tokens
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	Eof
)

func (t TokenType) String() string {
	switch t {
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"
	case Bang:
		return "Bang"
	case BangEqual:
		return "BangEqual"
	case Equal:
		return "Equal"
	case EqualEqual:
		return "EqualEqual"
	case Greater:
		return "Greater"
	case GreaterEqual:
		return "GreaterEqual"
	case Less:
		return "Less"
	case LessEqual:
		return "LessEqual"
	case Identifier:
		return "Identifier"
	case String:
		return "String"
	case Number:
		return "Number"
	case And:
		return "And"
	case Class:
		return "Class"
	case Else:
		return "Else"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case For:
		return "For"
	case If:
		return "If"
	case Nil:
		return "Nil"
	case Or:
		return "Or"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Super:
		return "Super"
	case This:
		return "This"
	case True:
		return "True"
	case Var:
		return "Var"
	case While:
		return "While"
	case Eof:
		return "Eof"
	default:
		return fmt.Sprintf("%d", int(t))
	}
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
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}
