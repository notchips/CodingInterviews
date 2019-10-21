package code

func ReverseSentence(bs []byte) {
	if len(bs) == 0 {
		return
	}
	indexes := []int{-1}
	for i, b := range bs {
		if b == ' ' {
			indexes = append(indexes, i)
		}
	}
	indexes = append(indexes, len(bs))
	// 翻转每个单词
	for i := 0; i < len(indexes)-1; i++ {
		l, r := indexes[i], indexes[i+1]
		reverseByteSlice(bs[l+1 : r])
	}
	// 翻转整个句子
	reverseByteSlice(bs)
}

func reverseByteSlice(bs []byte) {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}
