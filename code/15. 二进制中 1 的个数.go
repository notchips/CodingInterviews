func NumberOf1(n int) int {
	cnt := 0
	for ; n != 0; cnt++ {
		n &= n - 1
	}
	return cnt
}