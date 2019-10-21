package code

func NumDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, len(s))
	if justify(s[0:1]) { // 第一个字符能编码
		dp[0] = 1
	}
	for i := 1; i < n; i++ {
		if justify(s[i : i+1]) { // 自身能编码
			if justify(s[i-1 : i+1]) { // 能和前一个字符组合编码
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-1] + dp[i-2]
				}
			} else {
				dp[i] = dp[i-1]
			}
		} else { // 自身不能编码
			if justify(s[i-1 : i+1]) { // 能和前一个字符组合编码
				if i == 1 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-2]
				}
			} else {
				dp[i] = 0
			}
		}
	}
	return dp[n-1]
}

func justify(s string) bool {
	if len(s) == 1 && "1" <= s && s <= "9" {
		return true
	}
	if len(s) == 2 && "10" <= s && s <= "26" {
		return true
	}
	return false
}
