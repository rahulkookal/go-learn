package solutions

import "testing"

func TestKthLN(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
		name     string
	}{
		{
			arr:      []int{},
			k:        0,
			expected: 0,
			name:     "Empty Array",
		},
		{
			arr:      []int{3, 2, 6, 8},
			k:        2,
			expected: 6,
			name:     "Regular array and K",
		},
		{
			arr:      []int{3, 2, 6, 8},
			k:        1,
			expected: 8,
			name:     "Regular array and first Largest",
		},
		{
			arr:      []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
			name:     "Array with duplicate numbers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := KThLargestNumber(test.arr, test.k)
			if result != test.expected {
				t.Errorf("Exected: %d, got: %d", test.expected, result)
			}
		})
	}
}
