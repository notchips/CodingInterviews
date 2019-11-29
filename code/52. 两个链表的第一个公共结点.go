package code

func FindFirstCommonNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	l1, l2 := head1, head2
	l1MeetTail, l2MeetTail := 0, 0
	for l1 != l2 {
		if l1.Next == nil {
			l1 = head2
			l1MeetTail++
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = head1
			l2MeetTail++
		} else {
			l2 = l2.Next
		}
		if l1MeetTail == 2 || l2MeetTail == 2 { // 没有交点
			return nil
		}
	}
	return l1
}
