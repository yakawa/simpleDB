package value

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	UNKNOWN Type = iota
	EOS
	INTEGER
	IDENT
	K_SELECT
	S_SEMICOLON
	S_PLUS
	S_MINUS
	S_ASTERISK
	S_SOLIDAS
	S_PERCENT
)

func (t Type) String() string {
	switch t {
	case UNKNOWN:
		return "Unknown Value Type"
	case INTEGER:
		return "Integer Value Type"
	case IDENT:
		return "Identifier Value Type"
	case K_SELECT:
		return "Keyword (SELECT) "
	default:
		return "Unknwon Type"
	}
}

type Value struct {
	Type    Type
	Integer int
}

func Convert(s string) (Value, error) {
	isKeyword, valueType := CheckKeyword(s)

	if isKeyword {
		return Value{Type: valueType}, nil
	}

	if checkDigit([]rune(s)[0]) {
		for _, ch := range []rune(s) {
			if !checkDigit(ch) {
				return Value{Type: UNKNOWN}, errors.New(fmt.Sprintf("Unknown Format: %s", s))
			}
		}
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return Value{Type: UNKNOWN}, errors.New(fmt.Sprintf("Unknown Format: %s", s))
		}
		return Value{Type: INTEGER, Integer: int(v)}, nil
	} else {
		return Value{Type: IDENT}, nil
	}
}

func CheckKeyword(s string) (bool, Type) {
	switch strings.ToUpper(s) {
	case "SELECT":
		return true, K_SELECT
	}
	return false, UNKNOWN
}

func checkDigit(ch rune) bool {
	if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
		return true
	}
	return false
}
