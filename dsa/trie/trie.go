package trie

import "fmt"

// TrieNode represents a single node in the Trie
type TrieNode struct {
	isEnd    bool
	children map[rune]*TrieNode
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

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if next, exists := node.children[ch]; !exists {
			return false
		} else {
			node = next
		}
	}
	return node.isEnd
}

// Traverse prints all words stored in the Trie
func (t *Trie) Traverse() {
	var result []string
	t.findAll(t.root, "", &result)
	fmt.Println(result)
}

// TraversePre finds all words with the given prefix
func (t *Trie) TraversePre(prefix string) []string {
	node := t.root
	var result []string

	// Navigate to the end of the prefix
	for _, ch := range prefix {
		if next, exists := node.children[ch]; !exists {
			return result // Prefix not found
		} else {
			node = next
		}
	}

	// Collect words with the given prefix
	t.findAll(node, prefix, &result)
	return result
}

// findAll recursively finds all words from a given TrieNode
func (t *Trie) findAll(node *TrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}
	for ch, child := range node.children {
		t.findAll(child, prefix+string(ch), result)
	}
}
