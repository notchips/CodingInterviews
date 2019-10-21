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

type IntHeap struct {
	vals []int
	less func(int, int) bool
}

func (h IntHeap) Len() int           { return len(h.vals) }
func (h IntHeap) Less(i, j int) bool { return h.less(i, j) }
func (h IntHeap) Swap(i, j int)      { h.vals[i], h.vals[j] = h.vals[j], h.vals[i] }
func (h IntHeap) Top() int {
	return h.vals[0]
}

func (h *IntHeap) Push(x interface{}) {
	h.vals = append(h.vals, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h.vals
	n := len(old)
	x := old[n-1]
	h.vals = old[0 : n-1]
	return x
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
