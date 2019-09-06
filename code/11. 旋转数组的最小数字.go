func minNumberInRotateArray(array []int) int {
	n := len(array)
	if n == 0 {
		return math.MaxInt64
	}
	if n == 1 {
		return array[0]
	}
	left, right := 0, n-1
	mid := (left + right) / 2
	if array[left] < array[mid] { // 左边递增
		return min(array[left], minNumberInRotateArray(array[mid+1:]))
	} else if array[mid] < array[right] { // 右边递增
		return min(array[mid], minNumberInRotateArray(array[:mid]))
	}
	// 无法判断
	return min(minNumberInRotateArray(array[:mid]), minNumberInRotateArray(array[mid:]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}