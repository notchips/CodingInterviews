package code

func FindContinuousSequence(sum int) [][]int {
	if sum <= 0 {
		return nil
	}
	var answers [][]int
	left, right := 1, 2
	curSum := 3
	for right < sum {
		if curSum < sum {
			right++
			curSum += right
		} else if curSum > sum {
			curSum -= left
			left++
		} else {
			answers = append(answers, getAnswer(left, right))
			right++
			curSum += right
		}
	}
	return answers
}

func getAnswer(left, right int) []int {
	ans := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		ans = append(ans, i)
	}
	return ans
}
