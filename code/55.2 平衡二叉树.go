package code

func IsBalanced(root *TreeNode) bool {
	balanced := true
	checkHeight(root, &balanced)
	return balanced
}

func checkHeight(root *TreeNode, balanced *bool) int {
	if root == nil {
		return 0
	}
	leftHeight := checkHeight(root.Left, balanced)
	rightHeight := checkHeight(root.Right, balanced)
	if AbsInt(rightHeight-leftHeight) > 1 {
		*balanced = false
	}
	return MaxInt(leftHeight, rightHeight) + 1
}
