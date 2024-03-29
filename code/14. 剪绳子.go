package code

// 整数i可以拆分为2和i-2，2*(i-2) = 2i-4，当i>=4时,2i-4>=i，
// 故对于任意一个大于3的整数i都可以拆分为2和i-2，乘积>=原值
// 因此因子只需要2或3（1是最差的），又由于2*2*2 < 3*3
// 综上最优结果为：因子2不超过两个，其它因子全为3
func IntegerBreak(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	}

	threeCnt, tail := n/3, n%3
	switch tail {
	case 1:
		return pow(3, threeCnt-1) * 4
	case 2:
		return pow(3, threeCnt) * 2
	default:
		return pow(3, threeCnt)
	}
}

func pow(a, b int) int {
	if b < 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	if b&1 == 0 {
		t := pow(a, b/2)
		return t * t
	}
	return a * pow(a, b-1)
}
