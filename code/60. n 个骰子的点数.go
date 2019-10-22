package code

func DicesSum(n int) [][]float64 {
	if n < 1 {
		return nil
	}
	const faces = 6
	dp := make([][]float64, n+1) // dp[i][j] 表示i个骰子扔出j的次数
	for i := range dp {
		dp[i] = make([]float64, n*faces+1)
	}
	for i := 1; i <= faces; i++ { // 设置第1个骰子边界值
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ { // 计算第2到n个骰子
		for j := i; j <= i*faces; j++ { // i个骰子扔出的值范围为i到6*i
			for k := 1; k <= faces && k < j; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	total := Power(6, n) // n个骰子扔出的总可能数
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] /= total
		}
	}
	return dp
}
