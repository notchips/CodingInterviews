package code

import (
	"fmt"
	"sort"
	"strconv"
)

func MinNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s1 := fmt.Sprintf("%d%d", nums[i], nums[j])
		s2 := fmt.Sprintf("%d%d", nums[j], nums[i])
		return s1 < s2
	})
	ans := ""
	for _, num := range nums {
		ans += strconv.Itoa(num)
	}
	return ans
}
