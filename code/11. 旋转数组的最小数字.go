package code

import "math"

func minNum(a []int) int {
	n := len(a)

	switch n {
	case 0:
		return math.MaxInt64
	case 1:
		return a[0]
	}

	if a[0] < a[n-1] {
		return a[0]
	}

	return MinInt(minNum(a[:n/2]), minNum(a[n/2:]))
}
