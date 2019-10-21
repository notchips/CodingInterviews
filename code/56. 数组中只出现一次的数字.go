package code

func FindNumsAppearOnce(nums []int) (int, int) {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	diff &= -diff // 取diff最后一位1
	a, b := 0, 0
	for _, num := range nums {
		// 将nums分为两组，a和b在不同组中，组中其他数字都出现了两次
		if num&diff == 1 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return a, b
}
