package solutions

import "math"

// Helper function to calculate the total hours needed with speed `k`
func hoursRequired(piles []int, k int) int {
	totalHours := 0
	for _, pile := range piles {
		totalHours += int(math.Ceil(float64(pile) / float64(k))) // ⌈pile[i] / k⌉
	}
	return totalHours
}

// Function to find the minimum `k`
func MinEatingSpeed(piles []int, h int) int {
	// Find the range of `k`
	left, right := 1, 0
	for _, pile := range piles {
		if pile > right {
			right = pile // Max pile size is the upper limit for `k`
		}
	}

	// Perform binary search
	for left < right {
		mid := (left + right) / 2
		if hoursRequired(piles, mid) <= h {
			// Try a smaller k
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
