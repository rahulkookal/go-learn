package solutions

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	tests := []struct {
		name     string
		products []string
		search   string
		expected [][]string
	}{
		{
			name:     "Random words",
			products: []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			expected: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
			search: "mouse",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SuggestedProducts(test.products, test.search)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}

}
