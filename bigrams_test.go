package nlputils

import "testing"

func TestToBigrams(t *testing.T) {
	tcs := []testcase{
		{
			str:      "hello world",
			expected: []string{"<START>-=-hello", "world-=-<END>"},
		},
		{
			str:      "hello world, my name is JOhn!",
			expected: []string{"<START>-=-hello", "world-=-my", "my-=-name", "name-=-is", "is-=-john", "john-=-<END>"},
		},
	}

	for _, tc := range tcs {
		bigrams := ToBigrams(tc.str)
		if !tc.isMatch(bigrams) {
			t.Fatalf("invalid output, expected\n%v\nand received\n%v", tc.expected, bigrams)
		}
	}

}

type testcase struct {
	str      string
	expected []string
}

func (t *testcase) isMatch(a []string) (ok bool) {
	if len(a) != len(t.expected) {
		return
	}

	for i, value := range t.expected {
		if a[i] != value {
			return
		}
	}

	return true
}
