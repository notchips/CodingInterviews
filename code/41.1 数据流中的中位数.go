package code

import "container/heap"

var (
	left  = new(IntHeap)
	right = new(IntHeap)
)

func init() {
	left.less = func(i int, j int) bool {
		return left.vals[i] < left.vals[j]
	}
	right.less = func(i int, j int) bool {
		return right.vals[i] > right.vals[j]
	}
}

func InsertInt(v int) {
	n := left.Len() + right.Len()
	if n&1 == 0 {
		heap.Push(left, v)
		heap.Push(right, heap.Pop(left))
	} else {
		heap.Push(right, v)
		heap.Push(left, heap.Pop(right))
	}
}

func GetMedian() float64 {
	n := left.Len() + right.Len()
	if n&1 == 0 {
		return float64(left.Top()+right.Top()) / 2
	}
	return float64(right.Top())
}
