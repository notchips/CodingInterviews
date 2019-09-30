type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Convert(root *TreeNode) *TreeNode {
	head, _ := convert(root)
	return head
}

// convert 返回一个二叉搜索树对应非递减双链表的头结点和尾结点
func convert(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	leftHead, leftTail := convert(root.Left)
	rightHead, rightTail := convert(root.Right)
	newHead, newTail := root, root
	if leftTail != nil {
		leftTail.Right, root.Left = root, leftTail
		newHead = leftHead
	}
	if rightHead != nil {
		root.Right, rightHead.Left = rightHead, root
		newTail = rightTail
	}
	return newHead, newTail
}