package csv

import "testing"

func TestSplitColumn(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
		err      bool
	}{
		{
			input:    "colA, colB",
			expected: []string{"colA", "colB"},
			err:      false,
		},
		{
			input:    "colA, \"colA colB\"",
			expected: []string{"colA", "colA colB"},
			err:      false,
		},
	}

	for tn, tc := range testCases {
		r, err := splitColumn(tc.input)
		if err != nil && tc.err == false {
			t.Fatalf("Unexpected Exception: %s", err)
		}

		if len(tc.expected) != len(r) {
			t.Fatalf("[%d] Length mismatch", tn)
		}

		for n, col := range tc.expected {
			if col != r[n] {
				t.Fatalf("[%d] header mismatch %s != %s", tn, col, r[n])
			}
		}
	}
}
