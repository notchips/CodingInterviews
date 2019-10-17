package code

/*
	f(0) = 0
	f(1) = 1
	f(2) = f(1) + f(0)
	...
	f(n) = f(n-1) + f(n-2)
*/

func Fibonacci(n int) int {
	pre, post := 0, 1
	for i := 0; i < n; i++ {
		pre, post = post, pre+post
	}
	return pre
}
