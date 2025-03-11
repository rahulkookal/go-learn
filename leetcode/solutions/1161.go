package solutions

func MaxLevelSum(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxSum, maxLevel, level := root.Val, 1, 1

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		if maxSum < sum {
			maxSum = sum
			maxLevel = level
		}
		level++

	}
	return maxLevel
}
