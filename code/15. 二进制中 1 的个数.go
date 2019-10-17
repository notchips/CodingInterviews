package code

func NumberOf1(n int) int {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt
}
