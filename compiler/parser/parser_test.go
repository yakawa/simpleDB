package parser

import (
	"reflect"
	"testing"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/token"
	"github.com/yakawa/simpleDB/common/value"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		sql      string
		tokens   token.Tokens
		expected *ast.AST
	}{
		{
			sql: "SELECT 1;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT 1 + 2;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_PLUS,
					Literal: "+",
				},
				{
					Type:    token.NUMBER,
					Literal: "2",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 2,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT (1 + 2) * 3;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.S_LPAREN,
					Literal: "(",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_PLUS,
					Literal: "+",
				},
				{
					Type:    token.NUMBER,
					Literal: "2",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 2,
					},
				},
				{
					Type:    token.S_RPAREN,
					Literal: ")",
				},
				{
					Type:    token.S_ASTERISK,
					Literal: "*",
				},
				{
					Type:    token.NUMBER,
					Literal: "3",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 3,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT -1;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.S_MINUS,
					Literal: "-",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT +1;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.S_PLUS,
					Literal: "+",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT ABS(-1);",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.IDENT,
					Literal: "ABS",
				},
				{
					Type:    token.S_LPAREN,
					Literal: "(",
				},
				{
					Type:    token.S_MINUS,
					Literal: "-",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_RPAREN,
					Literal: ")",
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT 1,2;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.NUMBER,
					Literal: "1",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 1,
					},
				},
				{
					Type:    token.S_COMMA,
					Literal: ",",
				},
				{
					Type:    token.NUMBER,
					Literal: "2",
					Value: value.Value{
						Type:    value.INTEGER,
						Integer: 2,
					},
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
		{
			sql: "SELECT colA FROM tbl1;",
			tokens: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type:    token.IDENT,
					Literal: "colA",
				},
				{
					Type:    token.K_FROM,
					Literal: "FROM",
				},
				{
					Type:    token.IDENT,
					Literal: "tbl1",
				},
				{
					Type:    token.S_SEMICOLON,
					Literal: ";",
				},
				{
					Type: token.EOS,
				},
			},
			expected: &ast.AST{
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
		},
	}

	for tn, tc := range testCases {
		p, err := Parse(tc.tokens)
		if err != nil {
			t.Fatalf("[%d] %s : error: %s", tn, tc.sql, err)
		}

		if !reflect.DeepEqual(p, tc.expected) {
			t.Fatalf("[%d] %s Parse Error", tn, tc.sql)
		}
	}
}
