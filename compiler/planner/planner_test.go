package planner

import (
	"testing"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/runtime/vm"
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
		{
			sql: "SELECT (1 + 2) * 3;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											BinaryOperation: &ast.BinaryOpe{
												Operator: ast.B_ASTERISK,
												Left: &ast.Expression{
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
												Right: &ast.Expression{
													Literal: &ast.Literal{
														Numeric: &ast.Numeric{
															Integral: 3,
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
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 3,
					},
				},
				{
					Operator: vm.MUL,
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT -1;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											UnaryOperation: &ast.UnaryOpe{
												Operator: ast.U_MINUS,
												Expr: &ast.Expression{
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
						Integral: -1,
					},
				},
				{
					Operator: vm.MUL,
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT +1;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											UnaryOperation: &ast.UnaryOpe{
												Operator: ast.U_PLUS,
												Expr: &ast.Expression{
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
			sql: "SELECT ABS(-1);",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											FunctionCall: &ast.FunctionCall{
												Name: "ABS",
												Args: []ast.Expression{
													{
														UnaryOperation: &ast.UnaryOpe{
															Operator: ast.U_MINUS,
															Expr: &ast.Expression{
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
						Integral: -1,
					},
				},
				{
					Operator: vm.MUL,
				},
				{
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 1,
					},
				},
				{
					Operator: vm.CALL,
					Operand1: vm.VMValue{
						Type:   vm.String,
						String: "ABS",
					},
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT ABS(1);",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											FunctionCall: &ast.FunctionCall{
												Name: "ABS",
												Args: []ast.Expression{
													{
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
						Integral: 1,
					},
				},
				{
					Operator: vm.CALL,
					Operand1: vm.VMValue{
						Type:   vm.String,
						String: "ABS",
					},
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT 1,2;",
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
									{
										Expression: &ast.Expression{
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
				{
					Operator: vm.PUSH,
					Operand1: vm.VMValue{
						Type:     vm.Integer,
						Integral: 2,
					},
				},
				{
					Operator: vm.STORE,
				},
			},
		},
		{
			sql: "SELECT colA FROM tbl1;",
			ast: ast.AST{
				SQL: []ast.SQL{
					{
						SELECTStatement: &ast.SELECTStatement{
							Select: &ast.SELECTClause{
								ResultColumns: []ast.ResultColumn{
									{
										Expression: &ast.Expression{
											Column: &ast.Column{
												Column: "colA",
											},
										},
									},
								},
							},
							From: &ast.FROMClause{
								Table: &ast.Table{
									Table: "tbl1",
								},
							},
						},
					},
				},
			},
			expected: []vm.VMCode{
				{
					Operator: vm.READ,
					Operand1: vm.VMValue{
						Type: vm.Table,
						Table: vm.VMTable{
							Table:  "tbl1",
							DB:     "_",
							Schema: "LOCAL",
						},
					},
				},
				{
					Operator: vm.FETCH,
					Operand1: vm.VMValue{
						Type: vm.Column,
						Column: vm.VMColumn{
							Column: "colA",
							Table:  "tbl1",
							DB:     "_",
							Schema: "LOCAL",
						},
					},
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
			t.Fatalf("[%d] %s length mismatch %#+v", tn, tc.sql, vc)
		}
		for n, v := range vc {
			if v.Operator != tc.expected[n].Operator {
				t.Fatalf("[%d] %s OpCode mismatch", tn, tc.sql)
			}
		}
	}
}
