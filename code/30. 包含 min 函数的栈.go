package code

type stack struct {
	dataS []int
	minS  []int
	len   int
}

func NewStack(initCap int) *stack {
	return &stack{
		dataS: make([]int, 0, initCap),
		minS:  make([]int, 0, initCap),
		len:   0,
	}
}

func (s *stack) Push(val int) {
	s.dataS = append(s.dataS, val)
	if s.Empty() {
		s.minS = append(s.minS, val)
	} else {
		s.minS = append(s.minS, MinInt(val, s.Min()))
	}
	s.len++
}

func (s *stack) Pop() {
	if !s.Empty() {
		s.dataS = s.dataS[:s.len-1]
		s.minS = s.minS[:s.len-1]
		s.len--
	}
}

func (s *stack) Top() int {
	if !s.Empty() {
		return s.dataS[s.len-1]
	}
	return 0
}

func (s *stack) Min() int {
	if !s.Empty() {
		return s.minS[s.len-1]
	}
	return 0
}

func (s *stack) Empty() bool {
	return s.len == 0
}