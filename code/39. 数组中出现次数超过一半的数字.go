func MoreThanHalfNum(array []int) int {
	ans, cnt := 0, 0
	for _, num := range array {
		if cnt == 0 {
			ans, cnt = num, 1
		} else if ans == num {
			cnt++
		} else {
			cnt--
		}
	}
	for _, num := range array {
		if num == ans {
			cnt++
		}
	}
	if cnt > len(array)/2 {
		return ans
	}
	return 0
}