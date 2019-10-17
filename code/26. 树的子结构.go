package code

func HasSubtree(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	return isPartTreeFromRoot(root1, root2) ||
		HasSubtree(root1.Left, root2) || HasSubtree(root1.Right, root2)
}

func isPartTreeFromRoot(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil || root1.Val != root2.Val {
		return false
	}
	return isPartTreeFromRoot(root1.Left, root2.Left) &&
		isPartTreeFromRoot(root1.Right, root2.Right)
}