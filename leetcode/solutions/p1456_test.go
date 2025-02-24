package solutions

import "testing"

func TestMaxVowels(t *testing.T) {

	tests := []struct {
		name        string
		inputString string
		k           int
		expected    int
	}{
		{
			name:        "leetcode",
			inputString: "leetcode",
			k:           3,
			expected:    2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxVowels(test.inputString, test.k)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}

}
