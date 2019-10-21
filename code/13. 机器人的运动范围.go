package code

func MovingCount(threshold int, rows int, cols int) int {
	if rows <= 0 || cols <= 0 || threshold < 0 {
		return 0
	}
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	cnt := 0
	dfs(vis, 0, 0, threshold, &cnt)
	return cnt
}

func dfs(vis [][]bool, row, col, threshold int, cnt *int) {
	if row < 0 || row >= len(vis) ||
		col < 0 || col >= len(vis[0]) ||
		digitSum(row)+digitSum(col) > threshold ||
		vis[row][col] {
		return
	}
	vis[row][col] = true
	*cnt++
	dfs(vis, row+1, col, threshold, cnt)
	dfs(vis, row-1, col, threshold, cnt)
	dfs(vis, row, col+1, threshold, cnt)
	dfs(vis, row, col-1, threshold, cnt)
}

func digitSum(num int) int {
	sum := 0
	for num >= 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
