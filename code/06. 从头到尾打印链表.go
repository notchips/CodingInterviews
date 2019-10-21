package code

func ReverseListVal(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(ReverseListVal(head.Next), head.Val)
}
