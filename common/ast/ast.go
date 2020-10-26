package ast

type AST struct {
	SQL []SQL
}

type SQL struct {
	SELECTStatement *SELECTStatement
}

type SELECTStatement struct {
	Select *SELECTClause
	From   *FROMClause
}

type SELECTClause struct {
	ResultColumns []ResultColumn
}

type FROMClause struct {
	Table *Table
}

type ResultColumn struct {
	Expression *Expression
}

type Expression struct {
	Literal         *Literal
	UnaryOperation  *UnaryOpe
	BinaryOperation *BinaryOpe
	FunctionCall    *FunctionCall
	Column          *Column
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

	U_PLUS
	U_MINUS
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

	case U_PLUS:
		return "+"
	case U_MINUS:
		return "-"
	default:
		return "Unknwon Operation"
	}
}

type BinaryOpe struct {
	Operator OperatorType
	Left     *Expression
	Right    *Expression
}

type UnaryOpe struct {
	Operator OperatorType
	Expr     *Expression
}

type FunctionCall struct {
	Name string
	Args []Expression
}

type Column struct {
	Column string
	Table  string
	DB     string
	Schema string
}

type Table struct {
	Table  string
	DB     string
	Schema string
	Alias  string
}
