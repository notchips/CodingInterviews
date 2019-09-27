type ListNode struct {
	Val  int
	Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow := moveKStep(pHead, 1)
	fast := moveKStep(pHead, 2)
	for slow != fast {
		slow = moveKStep(slow, 1)
		fast = moveKStep(fast, 2)
	}
	fast = pHead
	for slow != fast {
		slow = moveKStep(slow, 1)
		fast = moveKStep(fast, 1)
	}
	return slow
}

func moveKStep(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}