package code

func isMatch(str string, pattern string) bool {
	str, pattern = str+" ", pattern+" "
	m, n := len(pattern), len(str)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for row := 2; row < m; row++ {
		if pattern[row] == '*' {
			dp[row][0] = dp[row-2][0]
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if pattern[row] == '.' || pattern[row] == str[col] {
				dp[row][col] = dp[row-1][col-1]
			} else if pattern[row] == '*' {
				if pattern[row-1] != '.' && pattern[row-1] != str[col] { // 只匹配到0个
					dp[row][col] = dp[row-2][col]
				} else { // 匹配0个 || 匹配1个 || 匹配多个
					dp[row][col] = dp[row-2][col] || dp[row-1][col] || dp[row][col-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
