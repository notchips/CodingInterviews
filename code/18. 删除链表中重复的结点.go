package code

func deleteDuplication(head *ListNode) *ListNode {
	headNode := new(ListNode)
	tempNode := headNode
	var pre *ListNode
	for head != nil {
		if !equal(pre, head) && !equal(head, head.Next) {
			tempNode.Next = head
			tempNode = tempNode.Next
		}
		pre, head = head, head.Next
	}
	tempNode.Next = nil
	return headNode.Next
}

func equal(a, b *ListNode) bool {
	if a != nil && b != nil {
		return a.Val == b.Val
	}
	return true
}
