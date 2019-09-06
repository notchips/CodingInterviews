func movingCount(threshold int, rows int, cols int) int {
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	cnt := 0
	dfs(threshold, rows, cols, 0, 0, vis, &cnt)
	return cnt
}

func dfs(threshold, rows, cols, row, col int, vis [][]bool, cnt *int) {
	if row < 0 || row >= rows ||
		col < 0 || col >= cols ||
		vis[row][col] {
		return
	}
	vis[row][col] = true
	if digitSum(row)+digitSum(col) > threshold {
		return
	}
	*cnt++
	dfs(threshold, rows, cols, row-1, col, vis, cnt)
	dfs(threshold, rows, cols, row+1, col, vis, cnt)
	dfs(threshold, rows, cols, row, col-1, vis, cnt)
	dfs(threshold, rows, cols, row, col+1, vis, cnt)
}

func digitSum(num int) int {
	sum := 0
	for num >= 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}