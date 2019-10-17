package code

//递归
//func ReverseList(head *ListNode) *ListNode {
//	return reverse(nil, head)
//}
//
//func reverse(pre, cur *ListNode) *ListNode {
//	if cur == nil {
//		return pre
//	}
//	post := cur.Next
//	cur.Next = pre
//	return reverse(cur, post)
//}

//头插法
func ReverseList(head *ListNode) *ListNode {
	headNode := new(ListNode)
	for head != nil {
		next := head.Next
		head.Next = headNode.Next
		headNode.Next = head
		head = next
	}
	return headNode.Next
}