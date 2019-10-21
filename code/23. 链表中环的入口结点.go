package code

func EntryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	if fast.Next == nil || fast.Next.Next == nil { // 无环
		return nil
	}
	slow, fast = slow.Next, fast.Next.Next
	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil { // 无环
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	fast = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}