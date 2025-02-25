package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type p437 struct{}

func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	count := 0
	new(p437).traverse(root, []int{}, targetSum, &count)
	return count
}

func (t p437) traverse(node *TreeNode, pathSums []int, target int, count *int) {
	if node == nil {
		return
	}

	// Update path sums
	newPathSums := make([]int, len(pathSums))
	copy(newPathSums, pathSums)

	for i := range newPathSums {
		newPathSums[i] += node.Val
		if newPathSums[i] == target {
			(*count)++
		}
	}

	// Check the node itself as a standalone path
	if node.Val == target {
		(*count)++
	}

	// Append current node value to path sums
	newPathSums = append(newPathSums, node.Val)

	// Recur for left and right children
	new(p437).traverse(node.Left, newPathSums, target, count)
	new(p437).traverse(node.Right, newPathSums, target, count)
}
