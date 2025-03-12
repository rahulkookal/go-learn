package solutions

import "testing"

func TestPivotIndex(t *testing.T) {

	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "array",
			arr:      []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "array with single element",
			arr:      []int{1},
			expected: 0,
		},
		{
			name:     "Array with negative numbers",
			arr:      []int{-1, -1, -1, -1, -1, 0},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PivotIndex(test.arr)
			if result != test.expected {
				t.Errorf("Test Failed input %d result: %d expected %d", test.arr, result, test.expected)
			}
		})
	}

}
