package code

func LayerOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var ans []int
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		ans = append(ans, front.Val)
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
	}
	return ans
}
