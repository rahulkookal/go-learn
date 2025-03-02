package solutions

import (
	"sort"
	"strings"
)

// TrieNode represents a single node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie structure
type Trie struct {
	root *TrieNode
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// TraversePre finds all words in the Trie with the given prefix
func (t *Trie) TraversePre(prefix string) []string {
	node := t.root
	var result []string

	// Navigate to the end of the prefix
	for _, ch := range prefix {
		if next, exists := node.children[ch]; exists {
			node = next
		} else {
			return result // Prefix not found
		}
	}

	// Collect words with the given prefix
	t.findAll(node, prefix, &result)
	sort.Strings(result) // Ensure lexicographical order

	// Limit results to 3 words
	if len(result) > 3 {
		return result[:3]
	}
	return result
}

// findAll performs DFS to collect words in the Trie
func (t *Trie) findAll(node *TrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}
	for ch, child := range node.children {
		t.findAll(child, prefix+string(ch), result)
	}
}

// suggestedProducts returns lexicographically sorted word suggestions for each prefix of searchWord
func SuggestedProducts(products []string, searchWord string) [][]string {
	trie := NewTrie()

	// Insert all products into the Trie
	for _, product := range products {
		trie.Insert(product)
	}

	var result [][]string
	var prefix strings.Builder

	// Find suggestions for each prefix of searchWord
	for _, ch := range searchWord {
		prefix.WriteRune(ch)
		words := trie.TraversePre(prefix.String())
		result = append(result, words)
	}

	return result
}
