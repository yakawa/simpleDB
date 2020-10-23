package vm

import (
	"github.com/yakawa/simpleDB/common/result"
)

type OpeType int

const (
	_ OpeType = iota
	PUSH
	POP
	STORE
)

func (o OpeType) String() string {
	switch o {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	case STORE:
		return "STORE"
	default:
		return "Unknwo Operation"
	}
}

type ValueType int

const (
	_ ValueType = iota
	Nothing
	Integer
)

func (v ValueType) String() string {
	switch v {
	case Nothing:
		return "No Value"
	case Integer:
		return "Integer"
	default:
		return "Unknown"
	}
}

type VMValue struct {
	Type     ValueType
	Integral int
}

type VMCode struct {
	Operator OpeType
	Operand1 VMValue
}

type VMStackValue struct {
	Value VMValue
}

func Run(codes []VMCode) []result.Value {
	s := newStack()
	cols := []result.Value{}

	for _, code := range codes {
		switch code.Operator {
		case PUSH:
			s.push(code.Operand1)
		case STORE:
			v, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			if v.Type == Integer {
				cols = append(cols, result.Value{Type: result.Integral, Integral: v.Integral})
			}
		}
	}
	return cols
}
