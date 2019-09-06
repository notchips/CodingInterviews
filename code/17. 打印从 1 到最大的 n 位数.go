func printOneToMaxNDigitNum(n int) {
	num := make([]byte, n)
	dfs(num, 0)
}

func dfs(num []byte, digit int) {
	if digit == len(num) {
		return
	}
	for i := 0; i <= 9; i++ {
		num[digit] = byte(i) + '0'
		printlnNum(num)
		dfs(num, digit+1)
	}
}

func printlnNum(num []byte) {
	for i, c := range num {
		if c != '0' {
			fmt.Println(string(num[i:]))
			return
		}
	}
}