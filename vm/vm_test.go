package vm

import (
	"testing"

	"github.com/yakawa/simpleDB/common/result"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		sql      string
		vmc      []VMCode
		expected []result.Value
	}{
		{
			sql: "SEELCT 1;",
			vmc: []VMCode{
				{
					Operator: PUSH,
					Operand1: VMValue{
						Type:     Integer,
						Integral: 1,
					},
				},
				{
					Operator: STORE,
				},
			},
			expected: []result.Value{
				{
					Type:     result.Integral,
					Integral: 1,
				},
			},
		},
		{
			sql: "SEELCT 1 + 2;",
			vmc: []VMCode{
				{
					Operator: PUSH,
					Operand1: VMValue{
						Type:     Integer,
						Integral: 1,
					},
				},
				{
					Operator: PUSH,
					Operand1: VMValue{
						Type:     Integer,
						Integral: 2,
					},
				},
				{
					Operator: ADD,
				},
				{
					Operator: STORE,
				},
			},
			expected: []result.Value{
				{
					Type:     result.Integral,
					Integral: 3,
				},
			},
		},
	}

	for tn, tc := range testCases {
		rslt := Run(tc.vmc)
		t.Logf("[%d] %#+v", tn, rslt)
		if len(tc.expected) != len(rslt) {
			t.Fatalf("[%d] %s Mistmach Result numbers", tn, tc.sql)
		}
		for n, r := range rslt {
			if r.Type != tc.expected[n].Type {
				t.Fatalf("[%d] %s Mistmach Result Type", tn, tc.sql)
			}
			switch r.Type {
			case result.Integral:
				if r.Integral != tc.expected[n].Integral {
					t.Fatalf("[%d] %s Mistmach Result", tn, tc.sql)
				}
			}
		}
	}
}
