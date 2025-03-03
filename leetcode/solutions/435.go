package solutions

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	pointer := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if pointer > intervals[i][0] {
			count++
		} else {
			pointer = intervals[i][1]
		}
	}

	return count
}
