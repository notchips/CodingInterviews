package code

func GetLeastNumbers(input []int, k int) []int {
	if len(input) < k {
		return nil
	}
	// 对前k个数建立大顶堆
	heap := input[:k]
	for i := len(heap) / 2; i >= 0; i-- {
		adjust(heap, i, len(heap)-1)
	}

	// 后k个数和堆顶作比较，小于则加入堆中
	for i := k; i < len(input); i++ {
		if input[i] < heap[0] {
			heap[0], input[i] = input[i], heap[0]
			adjust(heap, 0, len(heap)-1)
		}
	}
	return heap
}

func adjust(a []int, low, high int) {
	i, j := low, low*2+1
	for j <= high {
		if j+1 <= high && a[j+1] > a[j] {
			j++
		}
		if a[j] > a[i] {
			a[i], a[j] = a[j], a[i]
			i, j = j, j*2+1
		} else {
			return
		}
	}
}
