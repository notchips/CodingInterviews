package code

func Find(a [][]int, target int) bool {
	if len(a) == 0 || len(a[0]) == 0 {
		return false
	}
	m, n := len(a), len(a[0])
	row, col := 0, n-1 //从右上角开始遍历，每次可排除一行或一列
	for row < m-1 && col > 0 {
		if a[row][col] == target {
			return true
		} else if a[row][col] < target {
			row++
		} else {
			col--
		}
	}
	if row == m-1 { // 遍历到最后一行，二分搜索
		left, right := 0, col
		for left <= right {
			mid := left + (right-left)/2
			if a[m-1][mid] == target {
				return true
			} else if a[m-1][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else { // col == 0， 遍历到第一列，二分搜索
		top, down := row, m-1
		for top <= down {
			mid := top + (down-top)/2
			if a[mid][0] == target {
				return true
			} else if a[mid][0] < target {
				top = mid + 1
			} else {
				down = mid - 1
			}
		}
	}
	return false
}
