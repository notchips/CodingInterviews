package code

import "strconv"

/*
0-9，10个1位数字，范围：0-9
10-99，90个2位数字，范围：10-189
100-999，900个3位数字，范围：190-2899
1000-9999，9000个4位数字，范围：2900-38899
*/

func GetDigitAtIndex(n int) int {
	if n < 0 {
		return 0
	}
	var ans int
	for digit := 1; n >= 0; digit++ {
		numCnt := getNumCnt(digit)
		if n < numCnt*digit {
			firstNum := getFirstNum(digit)
			curNum := firstNum + n/digit
			curChar := strconv.Itoa(curNum)[n%digit]
			ans = int(curChar - '0')
		}
		n -= numCnt * digit
	}
	return ans
}

// 10, 90, 900, 9000...
func getNumCnt(digit int) int {
	if digit == 1 {
		return 10
	}
	return pow(10, digit-1) * 9
}

// 0, 10, 100, 1000...
func getFirstNum(digit int) int {
	if digit == 1 {
		return 0
	}
	return pow(10, digit-1)
}
