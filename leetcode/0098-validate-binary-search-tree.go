package leetcode

func isValidBST(root *TreeNode) bool {
	var stack TreeNodeStack
	var last *TreeNode
	current := root
	for current != nil || !stack.IsEmpty() {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		current = stack.Pop()
		if last != nil && last.Val >= current.Val {
			return false
		}

		last, current = current, current.Right
	}
	return true
}

type TreeNodeStack []*TreeNode

func (s *TreeNodeStack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *TreeNodeStack) Pop() *TreeNode {
	last := len(*s) - 1
	node := (*s)[last]
	*s = (*s)[:last]
	return node
}

func (s *TreeNodeStack) IsEmpty() bool {
	return len(*s) == 0
}
