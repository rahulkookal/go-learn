package solutions

import "testing"

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "array with integers",
			input:    []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "hight zero",
			input:    []int{0, -4, -7, -9, -10, -6, -3, -1},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LargestAltitude(test.input)
			if result != test.expected {
				t.Errorf("Test failed. Result: %d expected %d", result, test.expected)
			}
		})
	}
}
