func duplicate(numbers []int, length int, duplication *int) bool {
	for i := range numbers {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				*duplication = numbers[i]
				return true
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return false
}