package code

import "sort"

func VerifyPostOrder(sequence []int) bool {
	clone := make([]int, len(sequence))
	copy(clone, sequence)
	sort.Ints(sequence)
	return judge(sequence, clone)
}

func judge(in, post []int) bool {
	if len(in) == 0 {
		return true
	}
	mid := -1
	for i, v := range in {
		if v == post[len(post)-1] {
			mid = i
			break
		}
	}
	if mid == -1 {
		return false
	}
	return judge(in[:mid], post[:mid]) &&
		judge(in[mid+1:], post[mid:len(post)-1])
}