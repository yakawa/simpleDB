package token

import (
	"strings"

	"github.com/yakawa/simpleDB/common/value"
)

type Type int

const (
	UNKNOWN Type = -1
	ERROR   Type = iota
	EOS
	IDENT
	NUMBER

	K_SELECT
	K_FROM

	S_PLUS
	S_MINUS
	S_ASTERISK
	S_SOLIDAS
	S_PERCENT
	S_SEMICOLON
	S_LPAREN
	S_RPAREN
	S_COMMA
)

func (t Type) String() string {
	switch t {
	case UNKNOWN:
		return "Unknown Token"
	case ERROR:
		return "Error Token"
	case EOS:
		return "End of Sentence Token"
	case IDENT:
		return "Identifier Token"
	case NUMBER:
		return "Number Token"

	case K_SELECT:
		return "Keyword (SELECT)"
	case K_FROM:
		return "Keyword (FROM)"

	case S_PLUS:
		return "Symbol (+)"
	case S_MINUS:
		return "Symbol (-)"
	case S_ASTERISK:
		return "Symbol (*)"
	case S_SOLIDAS:
		return "Symbol (/)"
	case S_PERCENT:
		return "Symbol (%)"
	case S_SEMICOLON:
		return "Symbol (;)"
	case S_LPAREN:
		return "Symbol (()"
	case S_RPAREN:
		return "Symbol ())"
	case S_COMMA:
		return "Symbol (,)"

	default:
		return "Unknown Type"
	}
}

type Token struct {
	Type    Type
	Literal string
	Value   value.Value
}

type Tokens []Token

func (t Tokens) GetN(n int) Token {
	if len(t) <= n {
		return Token{
			Type: ERROR,
		}
	}
	return t[n]
}

func CheckKeyword(s string) (bool, Type) {
	switch strings.ToUpper(s) {
	case "SELECT":
		return true, K_SELECT
	case "FROM":
		return true, K_FROM
	}
	return false, UNKNOWN
}
