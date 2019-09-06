func replaceSpace(str *[]byte) {
	oldLen := len(*str)
	spaceCnt := 0
	for _, c := range *str {
		if c == ' ' {
			spaceCnt++
		}
	}
	newLen := oldLen + 2*spaceCnt
	for i := 0; i < 2*spaceCnt; i++ {
		*str = append(*str, byte(0))
	}
	// p1指向原字符串的最后一个字节，p2指向扩充后的最后一个字节
	for p1, p2 := oldLen-1, newLen; p1 >= 0; p1-- {
		if (*str)[p1] == ' ' {
			(*str)[p2-2], (*str)[p2-1], (*str)[p2] = '%', '2', '0'
			p2 -= 3
		} else {
			(*str)[p2] = (*str)[p1]
			p2--
		}
	}
}