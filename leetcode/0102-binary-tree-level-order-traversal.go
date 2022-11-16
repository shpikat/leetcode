package leetcode

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int
	if root != nil {
		queue := []*TreeNode{root}
		for len(queue) != 0 {
			var current []int
			var next []*TreeNode
			for _, node := range queue {
				current = append(current, node.Val)
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
			queue = next
			levels = append(levels, current)
		}
	}
	return levels
}
