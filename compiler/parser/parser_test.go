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
					Type:    token.KEYWORD,
					Literal: "SELECT",
					Value: value.Value{
						Type: value.K_SELECT,
					},
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
					Type:    token.SYMBOL,
					Literal: ";",
					Value: value.Value{
						Type: value.S_SEMICOLON,
					},
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
					Type:    token.KEYWORD,
					Literal: "SELECT",
					Value: value.Value{
						Type: value.K_SELECT,
					},
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
					Type:    token.SYMBOL,
					Literal: "+",
					Value: value.Value{
						Type: value.S_PLUS,
					},
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
					Type:    token.SYMBOL,
					Literal: ";",
					Value: value.Value{
						Type: value.S_SEMICOLON,
					},
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
	}

	for tn, tc := range testCases {
		p, err := Parse(tc.tokens)
		if err != nil {
			t.Fatalf("[%d] %s : error: %s", tn, tc.sql, err)
		}

		if !reflect.DeepEqual(p, tc.expected) {
			t.Fatalf("[%d] %s Parse Error :%#+v:", tn, tc.sql, p.SQL[0].SELECTStatement.Select.ResultColumns[0])
		}
	}
}
