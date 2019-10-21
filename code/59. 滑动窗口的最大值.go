package code

import "container/list"

func MaxInWindows(nums []int, size int) []int {
	if len(nums) == 0 || size < 1 || size > len(nums) {
		return nil
	}
	if size == 1 {
		return nums
	}
	doubleList := list.New()
	var ans []int
	// 将前size个数的索引加入链表，保持链头索引对应数字最大
	for i := 0; i < size; i++ {
		for doubleList.Len() > 0 && nums[doubleList.Back().Value.(int)] <= nums[i] {
			doubleList.Remove(doubleList.Back())
		}
		doubleList.PushBack(i)
	}
	ans = append(ans, nums[doubleList.Front().Value.(int)])
	// 滑动窗口
	for l, r := 1, size; r < len(nums); l, r = l+1, r+1 {
		// 删除过期索引
		for doubleList.Len() > 0 && doubleList.Front().Value.(int) < l {
			doubleList.Remove(doubleList.Front())
		}
		// 保持链头索引对应数字最大
		for doubleList.Len() > 0 && nums[doubleList.Back().Value.(int)] <= nums[r] {
			doubleList.Remove(doubleList.Back())
		}
		doubleList.PushBack(r)
		ans = append(ans, nums[doubleList.Front().Value.(int)])
	}
	return ans
}
