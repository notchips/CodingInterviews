package code

func InversePairs(a []int) int {
	cnt := 0
	n := len(a)
	for step := 2; step/2 < n; step *= 2 {
		for left := 0; left < n; left += step {
			mid := left + step/2 - 1
			if mid < n-1 {
				right := MinInt(n-1, left+step-1)
				merge(a, left, mid, right, &cnt)
			}
		}
	}
	return cnt
}

func merge(a []int, left, mid, right int, cnt *int) {
	i, j, k := left, mid+1, 0
	b := make([]int, right-left+1)
	for i <= mid && j <= right {
		if a[i] <= a[j] {
			b[k] = a[i]
			i, k = i+1, k+1
		} else {
			b[k] = a[j]
			j, k = j+1, k+1
			*cnt += mid - i + 1
		}
	}
	for i <= mid {
		b[k] = a[i]
		i, k = i+1, k+1
	}
	for j <= right {
		b[k] = a[j]
		j, k = j+1, k+1
	}
	for i, k = left, 0; i <= right; i, k = i+1, k+1 {
		a[i] = b[k]
	}
}
