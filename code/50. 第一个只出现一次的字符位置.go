package code

func FirstNotRepeatingChar(s string) (int, bool) {
	var charCnt [256]int
	for i := 0; i < len(s); i++ {
		charCnt[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if charCnt[s[i]] == 1 {
			return i, true
		}
	}
	return -1, false
}
