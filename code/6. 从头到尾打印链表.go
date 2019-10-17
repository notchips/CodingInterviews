package code

func reverseList(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reverseList(head.Next), head.Val)
}
