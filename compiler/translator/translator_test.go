package translator

import (
	"testing"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/vm"
)

func TestTranslate(t *testing.T) {
	testCases := []struct {
		sql      string
		ast      ast.AST
		expected []vm.VMCode
	}{
		{
			sql: "SELECT 1;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											Literal: &ast.Literal{
												Numeric: &ast.Numeric{
													Integral: 1,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []vm.VMCode{
				{
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 1,
					},
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT 1 + 2;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											BinaryOperation: &ast.BinaryOpe{
												Operator: ast.B_PLUS,
												Left: &ast.Expression{
													Literal: &ast.Literal{
														Numeric: &ast.Numeric{
															Integral: 1,
														},
													},
												},
												Right: &ast.Expression{
													Literal: &ast.Literal{
														Numeric: &ast.Numeric{
															Integral: 2,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []vm.VMCode{
				{
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 1,
					},
				},
				{
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 2,
					},
				},
				{
					Operator: vm.ADD,
				},
				{
					Operator: vm.STORE,
				},
			},
		},
	}

	for tn, tc := range testCases {
		vc := Translate(&tc.ast)
		if len(vc) != len(tc.expected) {
			t.Fatalf("[%d] %s length mismatch", tn, tc.sql)
		}
		for n, v := range vc {
			if v.Operator != tc.expected[n].Operator {
				t.Fatalf("[%d] %s OpCode mismatch", tn, tc.sql)
			}
		}
	}
}
