package code

func GetNumberOfK(nums []int, k int) int {
	l, r := findFirstKIndex(nums, k), findLastKIndex(nums, k)
	if l == -1 || r == -1 {
		return 0
	}
	return r - l + 1
}

func findFirstKIndex(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == k {
		return left
	}
	return -1
}

func findLastKIndex(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left+right)/2 + 1
		if nums[mid] > k {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[left] == k {
		return left
	}
	return -1
}
