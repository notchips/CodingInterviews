type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return isMirror(pRoot.Left, pRoot.Right)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}