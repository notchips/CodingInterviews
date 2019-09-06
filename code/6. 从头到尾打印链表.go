type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(printListFromTailToHead(head.Next), head.Val)
}