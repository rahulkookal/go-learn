package solutions

import "strings"

func RemoveStars(s string) string {
	newStr := s
	for strings.Contains(newStr, "*") {
		i := strings.Index(newStr, "*")
		if i > 0 {
			newStr = newStr[:i-1] + newStr[i+1:] // Remove '*' and the character before it
		} else {
			newStr = newStr[i+1:] // If '*' is at index 0, remove only '*'
		}
	}
	return newStr
}
