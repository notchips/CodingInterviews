package code

import "math"

/*
	f(0) = 0
	f(1) = 1
	f(2) = f(1) + 1 = 2f(1) = 2
	f(3) = f(2) + f(1) + 1 = 4f(1) = 4
	f(4) = f(3) + f(2) + f(1) + 1 = 8f(1) = 8
*/
func JumpFloorII(target int) int {
	if target <= 0 {
		return 0
	}
	return int(math.Pow(2, float64(target-1)))
}
