package solutions

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]  // Corrected node extraction
			queue = queue[1:] // Dequeue

			// Enqueue left and right children
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// Add the last node at this level to the result
			if i == n-1 {
				result = append(result, node.Val)
			}
		}
	}

	return result
}
