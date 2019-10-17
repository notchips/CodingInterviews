package code

//递归
//func Merge(list1, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//	if list2 == nil {
//		return list1
//	}
//	if list1.Val < list2.Val {
//		list1.Next = Merge(list1.Next, list2)
//		return list1
//	}
//	list2.Next = Merge(list1, list2.Next)
//	return list2
//}

//迭代
func Merge(list1, list2 *ListNode) *ListNode {
	headNode := new(ListNode)
	preNode := headNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			preNode.Next = list1
			list1 = list1.Next
		} else {
			preNode.Next = list2
			list2 = list2.Next
		}
		preNode = preNode.Next
	}
	for list1 != nil {
		preNode.Next = list1
		list1 = list1.Next
	}
	for list2 != nil {
		preNode.Next = list2
		list2 = list2.Next
	}
	return headNode.Next
}