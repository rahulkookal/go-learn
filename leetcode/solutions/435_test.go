package solutions

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {

	tests := []struct {
		name     string
		expected int
		input    [][]int
	}{
		{
			name:     "No overlapping intervals",
			expected: 0,
			input:    [][]int{{1, 2}, {2, 3}},
		},
		{
			name:     "Single overlapping intervals",
			expected: 1,
			input:    [][]int{{1, 2}, {2, 3}, {2, 5}},
		},
		{
			name:     "Same intervals",
			expected: 2,
			input:    [][]int{{1, 2}, {1, 2}, {1, 2}},
		},
		{
			name:     "Empty intervals",
			expected: 0,
			input:    [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EraseOverlapIntervals(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}

}
