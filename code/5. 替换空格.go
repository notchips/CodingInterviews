func replaceSpace(str *[]byte) {
	p1 := len(*str) - 1
	for i := 0; i <= p1; i++ {
		if (*str)[i] == ' ' {
			*str = append(*str, byte(0), byte(0))
		}
	}
	for p2 := len(*str) - 1; p1 >= 0; p1-- {
		if (*str)[p1] == ' ' {
			(*str)[p2-2], (*str)[p2-1], (*str)[p2] = '%', '2', '0'
			p2 -= 3
		} else {
			(*str)[p2] = (*str)[p1]
			p2--
		}
	}
}