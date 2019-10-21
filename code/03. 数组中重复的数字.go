package code

func Duplicate(a []int) (int, bool) {
	for _, v := range a {
		if v < 0 || v >= len(a) {
			return 0, false
		}
	}
	for i := range a {
		for a[i] != i {
			if a[i] == a[a[i]] {
				return a[i], true
			}
			a[i], a[a[i]] = a[a[i]], a[i]
		}
	}
	return 0, false
}
