type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	var mid int
	for i, v := range in {
		if v == pre[0] {
			mid = i
			break
		}
	}
	return &TreeNode{
		Val:   pre[0],
		Left:  reConstructBinaryTree(pre[1:mid+1], in[:mid]),
		Right: reConstructBinaryTree(pre[mid+1:], in[mid+1:]),
	}
}