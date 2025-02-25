package solutions

import "strconv"

func Compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	var result []byte
	a := chars[0]
	i := 0

	for _, char := range chars {
		if char == a {
			i++
		} else {
			result = append(result, a) // Append previous character
			if i > 1 {
				result = append(result, []byte(strconv.Itoa(i))...) // Append count as bytes
			}
			a = char
			i = 1
		}
	}

	// Append the last character and count
	result = append(result, a)
	if i > 1 {
		result = append(result, []byte(strconv.Itoa(i))...)
	}

	// Copy back to `chars` to modify it in place
	copy(chars, result)

	return len(result)
}
