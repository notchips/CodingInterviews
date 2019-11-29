package code

func Add(a, b int) int {
	if b == 0 {
		return a
	}
	return Add(a^b, a&b<<1)
}
