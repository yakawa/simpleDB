package translator

import (
	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/vm"
)

func Translate(a *ast.AST) []vm.VMCode {
	codes := []vm.VMCode{}
	for _, sql := range a.SQL {
		for _, col := range sql.SELECTStatement.Select.ResultColumns {
			c := translateResultColumn(col)
			codes = append(codes, c)

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

func translateResultColumn(c ast.ResultColumn) vm.VMCode {
	code := vm.VMCode{}
	v := vm.VMValue{}
	if c.Expression.Literal != nil {
		if c.Expression.Literal.Numeric != nil {
			v.Type = vm.Integer
			v.Integral = c.Expression.Literal.Numeric.Integral
			code.Operator = vm.PUSH
			code.Operand1 = v
		}
	}
	return code
}
