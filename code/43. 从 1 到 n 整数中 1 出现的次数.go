package code

func NumberOf1Between1AndN(n int) int {
	cnt := 0
	for m := 1; m <= n; m *= 10 {
		pre, post := n/m, n%m
		if pre%10 == 0 {
			cnt += pre / 10 * m
		} else if pre%10 == 1 {
			cnt += pre/10*m + post + 1
		} else {
			cnt += pre/10*m + m
		}
	}
	return cnt
}
