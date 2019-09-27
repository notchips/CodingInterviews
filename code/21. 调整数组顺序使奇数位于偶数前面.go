//func reOrderArray(array []int) {
//	n := len(array)
//	for i := 1; i < n; i++ {
//		changed := false
//		for j := 0; j < n-i; j++ {
//			if array[j]&1 == 0 && array[j+1]&1 == 1 {
//				array[j], array[j+1] = array[j+1], array[j]
//				changed = true
//			}
//		}
//		if !changed {
//			return
//		}
//	}
//}

func reOrderArray(array []int) {
	clone := make([]int, len(array))
	copy(clone, array)
	oddCnt := 0
	for _, num := range array {
		if num&1 == 1 {
			oddCnt++
		}
	}
	oddPtr, evenPtr := 0, oddCnt
	for _, num := range clone {
		if num&1 == 1 {
			array[oddPtr] = num
			oddPtr++
		} else {
			array[evenPtr] = num
			evenPtr++
		}
	}
}