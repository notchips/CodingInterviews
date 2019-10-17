package code

import "errors"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("empty stack")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

type Queue struct {
	in  *Stack
	out *Stack
}

func (q *Queue) Push(v int) {
	q.in.Push(v)
}

func (q *Queue) Pop() (int, error) {
	if q.Empty() {
		return 0, errors.New("empty queue")
	}
	if q.out.Empty() {
		for !q.in.Empty() {
			v, _ := q.in.Pop()
			q.out.Push(v)
		}
	}
	return q.out.Pop()
}

func (q Queue) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}
