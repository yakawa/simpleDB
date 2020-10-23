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
	}

	for tn, tc := range testCases {
		p := Parse(tc.tokens)

		if !reflect.DeepEqual(p, tc.expected) {
			t.Fatalf("[%d] %s Parse Error", tn, tc.sql)
		}
	}
}
