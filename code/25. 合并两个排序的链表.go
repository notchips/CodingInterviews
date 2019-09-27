type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = Merge(list1.Next, list2)
		return list1
	}
	list2.Next = Merge(list1, list2.Next)
	return list2
}