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
					Type:    token.KEYWORD,
					Literal: "SELECT",
					Value: value.Value{
						Type: value.K_SELECT,
					},
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
					Value: value.Value{
						Type: value.IDENT,
					},
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
		},
		{
			input: "SELECT 1 - 1;",
			expected: token.Tokens{
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
					Literal: "-",
					Value: value.Value{
						Type: value.S_MINUS,
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
		},
	}

	for tn, tc := range testCases {
		tokens := Lex(tc.input)
		if len(tokens) != len(tc.expected) {
			t.Fatalf("[%d] expected tokens length %d, but got %d", tn, len(tc.expected), len(tokens))
		}
		for i, tk := range tokens {
			expected := tc.expected.GetN(i)
			if tk.Type != expected.Type {
				t.Fatalf("[%d] expected token type %s, but got %s", tn, expected.Type.String(), tk.Type.String())
			}
			if tk.Value.Type != expected.Value.Type {
				t.Fatalf("[%d] expected token type %s, but got %s", tn, expected.Value.Type.String(), tk.Value.Type.String())
			}
			if expected.Type == token.KEYWORD || expected.Type == token.IDENT {
				if tk.Literal != expected.Literal {
					t.Fatalf("[%d] expected token literal %s, but got %s", tn, expected.Literal, tk.Literal)
				}
			}
		}
	}
}
