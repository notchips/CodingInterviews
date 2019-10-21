package code

import "math"

func GetUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		next2, next3, next5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(next2, next3, next5)
		switch dp[i] {
		case next2:
			p2++
		case next3:
			p3++
		case next5:
			p5++
		}
	}
	return dp[n-1]
}

func min(a ...int) int {
	ret := math.MaxInt64
	for _, v := range a {
		if v < ret {
			ret = v
		}
	}
	return ret
}
