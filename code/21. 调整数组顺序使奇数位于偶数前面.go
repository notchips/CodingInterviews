package code

//BubbleSort
//func ReOrderArray(a []int) {
//	n := len(a)
//	for i := 1; i < n; i++ {
//		changed := false
//		for j := 0; j < n-i; j++ {
//			if a[j]&1 == 0 && a[j+1]&1 == 1 {
//				a[j], a[j+1] = a[j+1], a[j]
//				changed = true
//			}
//		}
//		if !changed {
//			return
//		}
//	}
//}

func ReOrderArray(a []int) {
	clone := make([]int, len(a))
	copy(clone, a)
	oddCnt := 0
	for _, num := range a {
		if num&1 == 1 {
			oddCnt++
		}
	}
	oddPtr, evenPtr := 0, oddCnt
	for _, num := range clone {
		if num&1 == 1 {
			a[oddPtr] = num
			oddPtr++
		} else {
			a[evenPtr] = num
			evenPtr++
		}
	}
}