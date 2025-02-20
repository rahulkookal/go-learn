package solutions

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
		target   int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
			target:   0,
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			target:   1,
			expected: 1,
		},
		{
			name: "Two-level tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			target:   3,
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
			target:   3,
			expected: 2,
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
			target:   3,
			expected: 2,
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
			target:   5,
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PathSum(test.root, test.target)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
