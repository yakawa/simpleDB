package lexer

import (
	"testing"

	"github.com/yakawa/simpleDB/common/token"
	"github.com/yakawa/simpleDB/common/value"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input    string
		expected token.Tokens
	}{
		{
			input: "SELECT",
			expected: token.Tokens{
				{
					Type:    token.K_SELECT,
					Literal: "SELECT",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			input: "MyDB",
			expected: token.Tokens{
				{
					Type:    token.IDENT,
					Literal: "MyDB",
				},
				{
					Type: token.EOS,
				},
			},
		},
		{
			input: "SELECT 1 + 2;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT 1 - 1;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT (1 + 2) * 3;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT -1;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT +1;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT ABS(-1);",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT 1,2;",
			expected: token.Tokens{
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
		},
		{
			input: "SELECT colA FROM tbl1;",
			expected: token.Tokens{
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
		},
	}

	for tn, tc := range testCases {
		tokens := Lex(tc.input)
		if len(tokens) != len(tc.expected) {
			t.Fatalf("[%d] %s expected tokens length %d, but got %d", tn, tc.input, len(tc.expected), len(tokens))
		}
		for i, tk := range tokens {
			expected := tc.expected.GetN(i)
			if tk.Type != expected.Type {
				t.Fatalf("[%d] %s expected token type %s, but got %s", tn, tc.input, expected.Type, tk.Type)
			}
			if expected.Type == token.IDENT {
				if tk.Literal != expected.Literal {
					t.Fatalf("[%d] expected token literal %s, but got %s", tn, expected.Literal, tk.Literal)
				}
			}
		}
	}
}
