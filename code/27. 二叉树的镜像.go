package code

func Mirror(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	Mirror(root.Left)
	Mirror(root.Right)
}