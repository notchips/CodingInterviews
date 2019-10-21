package code

func NumberOf1Between1AndN(n int) int {
	cnt := 0
	for m := 1; m <= n; m *= 10 {
		a, b := n/m, n%m
		if a%10 == 0 {
			cnt += a / 10 * m
		} else if a%10 == 1 {
			cnt += a/10*m + b + 1
		} else {
			cnt += a/10*m + m
		}
	}
	return cnt
}
