package code

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	slow, fast := pHead, pHead
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
	fast = pHead
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}