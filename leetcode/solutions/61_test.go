package solutions

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			m:        3,
			n:        7,
			name:     "3x7",
			expected: 28,
		},
		{
			m:        3,
			n:        2,
			name:     "3x2",
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UniquePaths(test.m, test.n)
			if result != test.expected {
				t.Errorf("Test failed for m:%d n:%d result: %d expected: %d", test.m, test.n, result, test.expected)
			}
		})
	}
}
