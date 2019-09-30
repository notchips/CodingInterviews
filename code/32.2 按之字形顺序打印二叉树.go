func PrintZigzagLevelOrder(root *TreeNode) []int {
	return bfs([]*TreeNode{root}, false)
}

func bfs(roots []*TreeNode, reverse bool) []int {
	if len(roots) == 0 {
		return nil
	}
	values := make([]int, 0, len(roots))
	children := make([]*TreeNode, 0, 2*len(roots))
	for _, root := range roots {
		values = append(values, root.Val)
		if root.Left != nil {
			children = append(children, root.Left)
		}
		if root.Right != nil {
			children = append(children, root.Right)
		}
	}
	if reverse {
		reverseInt(values)
	}
	reverse = !reverse
	return append(values, bfs(children, reverse)...)
}

func reverseInt(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}