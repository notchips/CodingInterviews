package code

func LengthOfLongestSubstring(s string) int {
	var charMap [256]int // charMap记录字符串s中字符对应的索引
	for i := range charMap {
		charMap[i] = -1
	}
	maxLen := 0
	// [left, right] 表示一个包含非重复字符的子串
	for left, right := 0, 0; right < len(s); right++ {
		if index := charMap[s[right]]; index >= left {
			left = index + 1
		}
		maxLen = MaxInt(maxLen, right-left+1)
		charMap[s[right]] = right
	}
	return maxLen
}
