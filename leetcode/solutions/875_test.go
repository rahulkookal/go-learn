package solutions

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		k        int
		expected int
	}{
		{
			name:     "Small Piles",
			piles:    []int{3, 6, 7, 11},
			k:        8,
			expected: 4,
		},
		{
			name:     "Large Piles",
			piles:    []int{30, 11, 23, 4, 20},
			k:        5,
			expected: 30,
		},
		{
			name:     "Very Large Piles",
			piles:    []int{100, 200, 300, 400, 500, 600, 700},
			k:        10,
			expected: 400,
		}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MinEatingSpeed(test.piles, test.k)
			if result != test.expected {
				t.Errorf("Piles: %d, k: %d, expected: %d result:%d", test.piles, test.k, test.expected, result)
			}
		})
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	tests := []struct {
		name  string
		piles []int
		k     int
	}{
		{
			name:  "Small Piles",
			piles: []int{3, 6, 7, 11},
			k:     8,
		},
		{
			name:  "Large Piles",
			piles: []int{30, 11, 23, 4, 20},
			k:     5,
		},
		{
			name:  "Very Large Piles",
			piles: []int{100, 200, 300, 400, 500, 600, 700},
			k:     10,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ { // `b.N` is controlled by Go’s benchmark runner
				MinEatingSpeed(test.piles, test.k)
			}
		})
	}
}
