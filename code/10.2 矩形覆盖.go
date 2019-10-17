package code

/*
	f(0) = 0
	f(1) = 1
	f(2) = 2
	f(3) = f(2) + f(1)
	...
	f(n) = f(n-1) + f(n-2)
*/

func RectCover(target int) int {
	if target <= 0 {
		return 0
	}
	pre, post := 1, 2
	for i := 0; i < target-1; i++ {
		pre, post = post, pre+post
	}
	return pre
}
