package code

func GetMost(board [][]int) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	m, n := len(board), len(board[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = board[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + board[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + board[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = MaxInt(dp[i-1][j], dp[i][j-1]) + board[i][j]
		}
	}
	return dp[m-1][n-1]
}
