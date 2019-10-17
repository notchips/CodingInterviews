package code

func ZigzagLayerOrder(root *TreeNode) []int {
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
		Reverse(values)
	}
	return append(values, bfs(children, !reverse)...)
}
