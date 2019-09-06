func deleteDuplication(head *ListNode) *ListNode {
	headNode := new(ListNode)
	tempNode := headNode
	for pre := (*ListNode)(nil); head != nil; head, pre = head.Next, head {
		if !equal(pre, head) && !equal(head, head.Next) {
			tempNode.Next = head
			tempNode = tempNode.Next
		}
	}
	tempNode.Next = nil
	return headNode.Next
}

func equal(a, b *ListNode) bool {
	return a != nil && b != nil && a.Val == b.Val
}