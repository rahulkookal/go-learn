package solutions

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
		name     string
	}{
		{
			[]int{1, 2, 2, 1, 1, 3},
			true,
			"Number with unique occurances",
		},
		{
			[]int{1, 2},
			false,
			"Two numbers with same occurance",
		},
		{
			[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			true,
			"Negative numbers with unique occurance",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UniqueOccurrences(test.input)
			if result != test.expected {
				t.Errorf("Test failed for %v", test.input)
			}
		})
	}
}
