package solutions

import "testing"

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		expected int
	}{
		{
			name:     "array with zeros",
			arr:      []int{1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1},
			k:        2,
			expected: 9,
		},
		{
			name:     "array starting with zeros",
			arr:      []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1},
			k:        2,
			expected: 7,
		},
		{
			name:     "array ending with zeros",
			arr:      []int{1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
			k:        2,
			expected: 7,
		},
		{
			name:     "array starting and ending with zeros",
			arr:      []int{0, 0, 1, 1, 1, 1, 0, 0, 0},
			k:        2,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LongestOnes(test.arr, test.k)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})

	}
}
