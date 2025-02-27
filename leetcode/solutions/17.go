package solutions

import "strconv"

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	digitToLetters := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	result := []string{}

	for _, digit := range digits {
		num, _ := strconv.Atoi(string(digit))
		letters := digitToLetters[num]

		if len(result) == 0 {
			result = append(result, letters...)
		} else {
			result = combine(result, letters)
		}
	}

	return result
}

func combine(existing []string, newLetters []string) []string {
	var res []string
	for _, prefix := range existing {
		for _, letter := range newLetters {
			res = append(res, prefix+letter)
		}
	}
	return res
}
