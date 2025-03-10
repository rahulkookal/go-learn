package solutions

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {

	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Two-level tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 3},
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
			expected: []int{1, 2, 3, 4},
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
			expected: []int{1, 2, 3, 4},
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
			expected: []int{1, 3, 7},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RightSideView(test.root)
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
