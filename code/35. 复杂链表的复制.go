package code

func Clone(head *RandomListNode) *RandomListNode {
	// 每个结点后插入一个相同label的结点
	for cur := head; cur != nil; {
		clone := &RandomListNode{Val: cur.Val}
		cur.Next, clone.Next = clone, cur.Next
		cur = clone.Next
	}
	// 令clone结点的random指针，指向clone的random结点
	for cur := head; cur != nil; {
		clone := cur.Next
		if cur.Random != nil {
			clone.Random = cur.Random.Next
		}
		cur = clone.Next
	}
	// 将clone的结点与原链表分离
	headNode := new(RandomListNode)
	for cur := head; cur != nil; {
		clone := cur.Next
		headNode.Next = clone
		cur, headNode = clone.Next, headNode.Next
	}
	return headNode.Next
}