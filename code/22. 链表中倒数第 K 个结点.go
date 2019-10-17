package code

func FindKthToTail(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}
	slow, fast := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return nil
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
