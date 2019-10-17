package code

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeLinkNode struct {
	Left   *TreeLinkNode
	Right  *TreeLinkNode
	Father *TreeLinkNode
}

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}