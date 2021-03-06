package planner

import (
	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/runtime/vm"
)

func Translate(a *ast.AST) []vm.VMCode {
	codes := []vm.VMCode{}
	for _, sql := range a.SQL {
		if sql.SELECTStatement.From != nil {
			c := translateFROM(sql.SELECTStatement.From)
			codes = append(codes, c...)
		}
		for _, col := range sql.SELECTStatement.Select.ResultColumns {
			c := translateResultColumn(col)
			codes = append(codes, c...)

			s := vm.VMCode{
				Operator: vm.STORE,
				Operand1: vm.VMValue{
					Type: vm.Nothing,
				},
			}
			codes = append(codes, s)
		}
	}
	return codes
}

func translateResultColumn(c ast.ResultColumn) []vm.VMCode {
	codes := translateExpression(c.Expression)
	return codes
}

func translateExpression(expr *ast.Expression) []vm.VMCode {
	codes := []vm.VMCode{}
	v := vm.VMValue{}
	if expr.Literal != nil {
		if expr.Literal.Numeric != nil {
			v.Type = vm.Integer
			v.Integral = expr.Literal.Numeric.Integral
			c := vm.VMCode{
				Operator: vm.PUSH,
				Operand1: v,
			}
			codes = append(codes, c)
		}
		return codes
	} else if expr.BinaryOperation != nil {
		cl := translateExpression(expr.BinaryOperation.Left)
		codes = append(codes, cl...)
		cr := translateExpression(expr.BinaryOperation.Right)
		codes = append(codes, cr...)

		var c vm.VMCode
		switch expr.BinaryOperation.Operator {
		case ast.B_PLUS:
			c = vm.VMCode{
				Operator: vm.ADD,
			}
		case ast.B_MINUS:
			c = vm.VMCode{
				Operator: vm.SUB,
			}
		case ast.B_ASTERISK:
			c = vm.VMCode{
				Operator: vm.MUL,
			}
		case ast.B_SOLIDAS:
			c = vm.VMCode{
				Operator: vm.DIV,
			}
		case ast.B_PERCENT:
			c = vm.VMCode{
				Operator: vm.MOD,
			}
		default:
			return codes
		}
		codes = append(codes, c)
		return codes
	} else if expr.UnaryOperation != nil {
		c := translateExpression(expr.UnaryOperation.Expr)
		codes = append(codes, c...)
		switch expr.UnaryOperation.Operator {
		case ast.U_MINUS:
			codes = append(codes, vm.VMCode{Operator: vm.PUSH, Operand1: vm.VMValue{Type: vm.Integer, Integral: -1}})
			codes = append(codes, vm.VMCode{Operator: vm.MUL})
		}
		return codes
	} else if expr.FunctionCall != nil {
		for _, arg := range expr.FunctionCall.Args {
			c := translateExpression(&arg)
			codes = append(codes, c...)
		}
		codes = append(codes, vm.VMCode{Operator: vm.PUSH, Operand1: vm.VMValue{Type: vm.Integer, Integral: len(expr.FunctionCall.Args)}})
		codes = append(codes, vm.VMCode{Operator: vm.CALL, Operand1: vm.VMValue{Type: vm.String, String: expr.FunctionCall.Name}})
		return codes
	} else if expr.Column != nil {
		c := vm.VMCode{
			Operator: vm.FETCH,
			Operand1: vm.VMValue{
				Column: vm.VMColumn{
					Column: expr.Column.Column,
					DB:     "_",
					Schema: "LOCAL",
				},
			},
		}
		codes = append(codes, c)
	}
	return codes
}

func translateFROM(from *ast.FROMClause) []vm.VMCode {
	codes := []vm.VMCode{}
	c := vm.VMCode{
		Operator: vm.READ,
		Operand1: vm.VMValue{
			Table: vm.VMTable{
				Table:  from.Table.Table,
				DB:     "_",
				Schema: "LOCAL",
			},
		},
	}
	codes = append(codes, c)
	return codes
}
