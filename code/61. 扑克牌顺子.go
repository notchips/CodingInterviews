package code

import "sort"

func isContinuous(nums []int) bool {
	if len(nums) == 0 || len(nums) > 13 {
		return false
	}
	sort.Ints(nums)
	zeroCnt := 0
	for _, num := range nums {
		if num == 0 {
			zeroCnt++
		} else if num < 0 && num > 13 {
			return false
		}
	}
	pos := 0
	for nums[pos] == 0 {
		pos++
	}
	spaceCnt := 0
	for ; pos < len(nums)-1; pos++ {
		diff := nums[pos+1] - nums[pos]
		if diff == 0 { // 存在重复牌
			return false
		}
		spaceCnt += diff
	}
	return zeroCnt >= spaceCnt
}
