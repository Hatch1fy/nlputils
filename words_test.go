package nlputils

import "testing"

func TestToWords(t *testing.T) {
	tcs := []testcase{
		{
			str:      "hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, tc := range tcs {
		words := ToWords(tc.str)
		if !tc.isMatch(words) {
			t.Fatalf("invalid output, expected %v and received %v", tc.expected, words)
		}
	}

}
