func Fibonacci(n int) int {
	pre, post := 0, 1
	for i := 0; i < n; i++ {
		pre, post = post, pre+post
	}
	return pre
}