package code

import "math"

func FindGreatestSumOfSubArray(a []int) int {
	ans := math.MinInt64
	sum := 0
	for _, v := range a {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
