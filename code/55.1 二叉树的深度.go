package code

func TreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MaxInt(TreeDepth(root.Left), TreeDepth(root.Right)) + 1
}
