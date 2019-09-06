func JumpFloor(target int) int {
	if target <= 0 {
		return 0
	}
	pre, post := 1, 1
	for i := 0; i < target; i++ {
		pre, post = post, pre+post
	}
	return pre
}