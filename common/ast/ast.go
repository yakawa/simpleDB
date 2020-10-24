package ast

type AST struct {
	SQL []SQL
}

type SQL struct {
	SELECTStatement *SELECTStatement
}

type SELECTStatement struct {
	Select *SELECTClause
}

type SELECTClause struct {
	ResultColumns []ResultColumn
}

type ResultColumn struct {
	Expression *Expression
}

type Expression struct {
	Literal         *Literal
	BinaryOperation *BinaryOpe
}

type Literal struct {
	Numeric *Numeric
}

type Numeric struct {
	Integral int
}

type OperatorType int

const (
	_ OperatorType = iota
	B_PLUS
	B_MINUS
	B_ASTERISK
	B_SOLIDAS
	B_PERCENT
)

func (o OperatorType) String() string {
	switch o {
	case B_PLUS:
		return "+"
	case B_MINUS:
		return "-"
	case B_ASTERISK:
		return "*"
	case B_SOLIDAS:
		return "/"
	case B_PERCENT:
		return "%"
	default:
		return "Unknwon Operation"
	}
}

type BinaryOpe struct {
	Operator OperatorType
	Left     *Expression
	Right    *Expression
}
