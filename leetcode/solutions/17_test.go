package solutions

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	tests := []struct {
		name     string
		expected []string
		input    string
	}{
		{
			"Random number 23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
			"23",
		},
		{
			"Empty number",
			[]string{},
			"",
		},
		{
			"Single digit",
			[]string{"a", "b", "c"},
			"2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LetterCombinations(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected: %v, got: %v", test.expected, result)
			}
		})
	}

}
