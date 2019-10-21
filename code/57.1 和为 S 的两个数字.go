package code

func FindNumbersWithSum(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return nums[left], nums[right]
		}
	}
	return 0, 0
}
