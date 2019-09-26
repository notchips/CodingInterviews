/*
	f(0) = 1
	f(1) = f(0) = 1
	f(2) = f(1) + f(0) = 2f(0) = 2
	f(3) = f(2) + f(1) + f(0) = 4f(0) = 4
	f(4) = f(3) + f(2) + f(1) + f(0) = 8f(0) = 8
*/
func JumpFloorII(target int) int {
	if target <= 0 {
		return 0
	}
	return int(math.Pow(2, float64(target-1)))
}