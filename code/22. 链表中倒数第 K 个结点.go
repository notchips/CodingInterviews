type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(head *ListNode, k int) *ListNode {
	pre, post := head, head
	for post != nil && k > 0 {
		post = post.Next
		k--
	}
	if k > 0 {
		return nil
	}
	for post != nil {
		pre, post = pre.Next, post.Next
	}
	return pre
}