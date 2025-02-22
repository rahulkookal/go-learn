package solutions

import "math/bits"

func CountBits(n int) []int {
	var result []int = []int{}
	for i := 0; i <= n; i++ {
		result = append(result, bits.OnesCount(uint(i)))
	}
	return result
}
