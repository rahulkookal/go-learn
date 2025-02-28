package solutions

func UniqueOccurrences(arr []int) bool {
	freq := map[int]int{}
	seen := map[int]bool{}
	for _, a := range arr {
		freq[a]++
	}
	for _, f := range freq {
		if seen[f] {
			return false
		}
		seen[f] = true
	}
	return true
}
