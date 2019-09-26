func minNumberInRotateArray(array []int) int {
	n := len(array)

	switch n {
	case 0:
		return math.MaxInt64
	case 1:
		return array[0]
	}

	if array[0] < array[n-1] {
		return array[0]
	}
	
	return minInt(minNumberInRotateArray(array[:n/2]), minNumberInRotateArray(array[n/2:]))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}