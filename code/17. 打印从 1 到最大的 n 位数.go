package code

import "fmt"

func printOneToNDigitMaxNum(n int) {
	num := make([]byte, n)
	setDigit(num, 0)
}

func setDigit(num []byte, digit int) {
	if digit == len(num) {
		printNum(num)
		return
	}
	for c := byte('0'); c <= '9'; c++ {
		num[digit] = c
		setDigit(num, digit+1)
	}
}

func printNum(num []byte) {
	for i, c := range num {
		if c != '0' {
			fmt.Println(string(num[i:]))
			return
		}
	}
}