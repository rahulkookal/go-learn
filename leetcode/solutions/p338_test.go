package solutions

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "Empty Array",
			input:    5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := CountBits(test.input)
			if !reflect.DeepEqual(out, test.expected) {
				t.Errorf("Exected: %v, got: %v", test.expected, out)
			}
		})
	}

}
