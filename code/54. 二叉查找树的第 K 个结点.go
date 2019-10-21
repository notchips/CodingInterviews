package code

func KthNode(root *TreeNode, k int) *TreeNode {
	if k < 1 {
		return nil
	}
	ch := make(chan *TreeNode)
	go inOrder(root, &k, ch)
	return <-ch
}

func inOrder(root *TreeNode, k *int, ch chan<- *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, k, ch)
	*k--
	if *k == 0 {
		ch <- root
	}
	inOrder(root.Right, k, ch)
}
