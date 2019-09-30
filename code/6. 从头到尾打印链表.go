type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) {
	if head == nil {
		return
	}
	printListFromTailToHead(head.Next)
	fmt.Println(head.Val)
}