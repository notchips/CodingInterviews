package code

func LeftRotateString(bs []byte, k int) {
	n := len(bs)
	reverseByteSlice(bs)
	reverseByteSlice(bs[:n-k])
	reverseByteSlice(bs[n-k:])
}
