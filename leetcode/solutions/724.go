package solutions

func PivotIndex(nums []int) int {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	leftSum := 0
	for i := 0; i < n; i++ {
		if leftSum == sum-nums[i] {
			return i
		}
		leftSum += nums[i]
		sum -= nums[i]
	}
	return -1
}
