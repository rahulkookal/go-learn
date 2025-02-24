package solutions

func MaxVowels(s string, k int) int {
	var out int = 0
	isVowel := map[byte]bool{
		'a': true,
		'e': true,
		'u': true,
		'i': true,
		'o': true,
	}

	for i := 0; i < k; i++ {
		if isVowel[s[i]] {
			out += 1
		}
	}
	var res int = out
	for j := 0; j < len(s)-k; j++ {
		if isVowel[s[j]] {
			out -= 1
		}
		if isVowel[s[j+k]] {
			out += 1
		}
		res = Max(out, res)
	}
	return res
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
