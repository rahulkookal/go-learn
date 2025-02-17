package solutions

import "testing"

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"a*b*c*", ""},
		{"abcdef", "abcdef"},
		{"**", ""},
		{"ab*c*de**fg", "afg"},
	}

	for _, test := range tests {
		result := RemoveStars(test.input)
		if result != test.expected {
			t.Errorf("RemoveStars(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
