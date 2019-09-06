func JumpFloorII(target int) int {
	ans := make([]int, target+1)
	sum := make([]int, target+1)
	for i := 1; i <= target; i++ {
		ans[i] = sum[i-1] + 1
		sum[i] = sum[i-1] + ans[i]
	}
	return ans[target]
}