package solutions

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "Two-level tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 2,
		},
		{
			name: "Unbalanced left-heavy tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: 4,
		},
		{
			name: "Unbalanced right-heavy tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			expected: 4,
		},
		{
			name: "Balanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxDepth(test.root)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
