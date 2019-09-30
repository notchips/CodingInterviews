func Permutation(str string) []string {
	var ans []string
	vis := make([]bool, len(str))
	buf := make([]byte, 0, len(str))
	dfs(str, vis, &buf, &ans)
	return ans
}

func dfs(str string, vis []bool, buf *[]byte, ans *[]string) {
	if len(*buf) == len(str) {
		newBuf := make([]byte, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, string(newBuf))
		return
	}
	for i := 0; i < len(str); i++ {
		if !vis[i] {
			vis[i] = true
			*buf = append(*buf, str[i])
			dfs(str, vis, buf, ans)
			*buf = (*buf)[:len(*buf)-1]
			vis[i] = false
		}
	}
}