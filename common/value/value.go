package value

import (
	"errors"
	"fmt"
	"strconv"
)

type Type int

const (
	UNKNOWN Type = iota
	EOS
	INTEGER
)

func (t Type) String() string {
	switch t {
	case UNKNOWN:
		return "Unknown Value Type"
	case EOS:
		return "End Of Sentence Type"
	case INTEGER:
		return "Integer Value Type"

	default:
		return "Unknwon Type"
	}
}

type Value struct {
	Type    Type
	Integer int
}

func Convert(s string) (Value, error) {
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
	}
	return Value{}, errors.New(fmt.Sprintf("Cloud not convert: %s", s))
}

func checkDigit(ch rune) bool {
	if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
		return true
	}
	return false
}
