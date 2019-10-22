package code

func LastRemaining(n, m int) int {
	if n <= 0 || m <= 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	return (LastRemaining(n-1, m) + m) % n
}
