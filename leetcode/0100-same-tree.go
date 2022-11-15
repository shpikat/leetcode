package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	} else if p == nil && q == nil {
		return true
	}
	return false
}
