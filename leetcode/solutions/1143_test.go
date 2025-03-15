package solutions

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {

	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "Same text",
			text1:    "rahul",
			text2:    "rahul",
			expected: 5,
		},
		{
			name:     "Common subsequence",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "empty",
			text1:    "",
			text2:    "",
			expected: 0,
		},
		{
			name:     "no common subsequence ",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LongestCommonSubsequence(test.text1, test.text2)
			if result != test.expected {
				t.Errorf(`Test failed %s:%s result: %d expected: %d`, test.text1, test.text2, result, test.expected)
			}
		})
	}

}
