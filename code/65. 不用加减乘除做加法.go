package code

func Add(a, b int) int {
	if b == 0 {
		return b
	}
	return Add(a^b, a&b<<1)
}
