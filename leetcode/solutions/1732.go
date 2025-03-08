package solutions

func LargestAltitude(gain []int) int {
	highest := 0
	current := 0
	for _, altitude := range gain {
		current += altitude
		if current > highest {
			highest = current
		}
	}
	return highest
}
