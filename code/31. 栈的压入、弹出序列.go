package code

func IsPopOrder(push, pop []int) bool {
	if len(push) != len(pop) {
		return false
	}
	stack := make([]int, 0, len(push))
	for i := 0; i < len(push); i++ {
		stack = append(stack, push[i])
		for len(pop) > 0 && pop[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			pop = pop[1:]
		}
	}
	return len(stack) == 0
}