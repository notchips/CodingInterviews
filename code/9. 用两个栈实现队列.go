type Queue struct {
	in  *Stack
	out *Stack
}

func (q Queue) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}

// Pop 在调用之前，必须保证Queue非空
func (q Queue) Pop() int {
	if q.out.Empty() {
		for !q.in.Empty() {
			q.out.Push(q.in.Pop())
		}
	}
	return q.out.Pop()
}

func (q Queue) Push(i int) {
	q.in.Push(i)
}

type Stack []int

func (s Stack) Empty() bool {
	return len(s) == 0
}

// Pop 在调用之前，必须保证Stack非空
func (s *Stack) Pop() int {
	n := len(*s)
	top := (*s)[n-1]
	*s = (*s)[:n-1]
	return top
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}