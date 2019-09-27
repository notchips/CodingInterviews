type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasSubtree(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	return IsPartTreeStartWithRoot(root1, root2) || HasSubtree(root1.Left, root2) || HasSubtree(root1.Right, root2)
}

func IsPartTreeStartWithRoot(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil || root1.Val != root2.Val {
		return false
	}
	return IsPartTreeStartWithRoot(root1.Left, root2.Left) && IsPartTreeStartWithRoot(root1.Right, root2.Right)
}