package solutions

func LongestOnes(nums []int, k int) int {
	left, result, zeroCount := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		if zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
