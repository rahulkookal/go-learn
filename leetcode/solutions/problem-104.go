package solutions

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// traverse(root, 1)
	return traverse(root, 1)
}

func traverse(node *TreeNode, depth int) int {

	if node == nil {
		return depth
	}

	depthLeft := depth
	depthRight := depth

	if node.Left != nil {
		depthLeft = traverse(node.Left, depth+1)
	}
	if node.Right != nil {
		depthRight = traverse(node.Right, depth+1)
	}

	return int(math.Max(float64(depthLeft), float64(depthRight)))
}
