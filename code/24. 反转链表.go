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
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = headNode.Next
		headNode.Next = cur
		cur = next
	}
	return headNode.Next
}