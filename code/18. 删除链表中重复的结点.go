func deleteDuplication(head *ListNode) *ListNode {
	headNode := new(ListNode)
	newPre := headNode
	pre := &ListNode{Next: head}
	for head != nil {
		if !equal(pre, head) && !equal(head, head.Next) {
			newPre.Next = head
			newPre = newPre.Next
		}
		head, pre = head.Next, head
	}
	newPre.Next = nil
	return headNode.Next
}

func equal(a, b *ListNode) bool {
	if a != nil && b != nil && a.Val == b.Val {
		return true
	}
	return a == b
}