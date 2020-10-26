package vm

import (
	"fmt"

	"github.com/yakawa/simpleDB/common/result"
	"github.com/yakawa/simpleDB/vm/functions"
)

type OpeType int

const (
	_ OpeType = iota
	PUSH
	POP
	ADD
	SUB
	MUL
	DIV
	MOD
	STORE
	CALL
	READ
	FETCH
)

func (o OpeType) String() string {
	switch o {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case MOD:
		return "MOD"
	case STORE:
		return "STORE"
	case CALL:
		return "CALL"
	case READ:
		return "READ"
	case FETCH:
		return "FETCH"
	default:
		return "Unknwo Operation"
	}
}

type ValueType int

const (
	_ ValueType = iota
	Nothing
	Integer
	String
	Table
	Column
)

func (v ValueType) String() string {
	switch v {
	case Nothing:
		return "No Value"
	case Integer:
		return "Integer"
	case String:
		return "String"
	case Table:
		return "Table"
	case Column:
		return "Column"
	default:
		return "Unknown"
	}
}

type VMValue struct {
	Type     ValueType
	Integral int
	String   string
	Table    VMTable
	Column   VMColumn
}

type VMTable struct {
	Table  string
	DB     string
	Schema string
}

type VMColumn struct {
	Column string
	Table  string
	DB     string
	Schema string
}
type VMCode struct {
	Operator OpeType
	Operand1 VMValue
	Operand2 VMValue
}

func (c VMCode) String() string {
	s := ""
	s = fmt.Sprintf("%s", c.Operator)

	if c.Operand1.Type != Nothing {
		switch c.Operand1.Type {
		case Integer:
			s = fmt.Sprintf("%s %d", s, c.Operand1.Integral)
		case String:
			s = fmt.Sprintf("%s %s", s, c.Operand1.String)
		}
	}

	if c.Operand2.Type != Nothing {
		switch c.Operand1.Type {
		case Integer:
			s = fmt.Sprintf("%s %d", s, c.Operand2.Integral)
		case String:
			s = fmt.Sprintf("%s %s", s, c.Operand2.String)
		}
	}

	return s
}

func Run(codes []VMCode) []result.Value {
	s := newStack()
	cols := []result.Value{}

	for _, code := range codes {
		switch code.Operator {
		case PUSH:
			s.push(code.Operand1)
		case ADD:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			v := VMValue{
				Type:     Integer,
				Integral: ope1.Integral + ope2.Integral,
			}
			s.push(v)

		case SUB:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			v := VMValue{
				Type:     Integer,
				Integral: ope1.Integral - ope2.Integral,
			}
			s.push(v)
		case MUL:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			v := VMValue{
				Type:     Integer,
				Integral: ope1.Integral * ope2.Integral,
			}
			s.push(v)

		case DIV:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			v := VMValue{
				Type:     Integer,
				Integral: ope1.Integral / ope2.Integral,
			}
			s.push(v)
		case MOD:
			ope2, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			ope1, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			v := VMValue{
				Type:     Integer,
				Integral: ope1.Integral % ope2.Integral,
			}
			s.push(v)

		case CALL:
			args := []interface{}{}

			argsN, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			for i := 0; i < argsN.Integral; i++ {
				v, err := s.pop()
				if err != nil {
					return []result.Value{}
				}
				switch v.Type {
				case Integer:
					args = append(args, v.Integral)
				case String:
					args = append(args, v.String)
				}
			}

			call := functions.LookupFunction(code.Operand1.String)
			if call == nil {
				return []result.Value{}
			}
			r := call(args)
			var vr VMValue
			switch r.Type {
			case result.Integral:
				vr = VMValue{
					Type:     Integer,
					Integral: r.Integral,
				}
			}
			s.push(vr)

		case STORE:
			v, err := s.pop()
			if err != nil {
				return []result.Value{}
			}
			if v.Type == Integer {
				cols = append(cols, result.Value{Type: result.Integral, Integral: v.Integral})
			}

		case READ:
		case FETCH:
		}
	}
	return cols
}
