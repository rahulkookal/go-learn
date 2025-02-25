package solutions

import (
	"testing"
)

func TestCompress(t *testing.T) {
	input1 := []string{"a", "a", "b", "b", "c", "c", "c"}
	input2 := []string{}
	input3 := []string{"a", "a", "b", "c", "c", "c"}

	type testCases struct {
		name     string
		input    []byte
		expected int
	}
	tcases := []testCases{
		{"String with more than one chars", convertStringToByteArray(input1), 6},
		{"String with Zero chars", convertStringToByteArray(input2), 0},
		{"String with single occurance of chars", convertStringToByteArray(input3), 5},
	}

	for _, test := range tcases {
		t.Run(test.name, func(t *testing.T) {
			result := Compress(test.input)
			if result != test.expected {
				t.Errorf("RemoveStars(%q) = %q; expected %q", test.input, result, test.expected)
			}
		})
	}

}

func convertStringToByteArray(in []string) []byte {
	o := make([]byte, len(in))
	for i, s := range in {
		o[i] = s[0]
	}
	return o
}
