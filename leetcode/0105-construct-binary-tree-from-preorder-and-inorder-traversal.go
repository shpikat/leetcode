package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	current := preorder[0]
	index := 0
	for inorder[index] != current {
		index++
	}
	return &TreeNode{
		current,
		buildTree(preorder[1:1+index], inorder[:index]),
		buildTree(preorder[index+1:], inorder[index+1:]),
	}
}
