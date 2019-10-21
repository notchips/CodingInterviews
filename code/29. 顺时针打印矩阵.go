package code

func PrintMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	rowStart, colStart := 0, 0
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1
	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			ans = append(ans, matrix[rowStart][i])
		}
		for j := rowStart + 1; j <= rowEnd-1; j++ {
			ans = append(ans, matrix[j][colEnd])
		}
		if rowEnd > rowStart {
			for i := colEnd; i >= colStart; i-- {
				ans = append(ans, matrix[rowEnd][i])
			}
		}
		if colEnd > colStart {
			for j := rowEnd - 1; j >= rowStart+1; j-- {
				ans = append(ans, matrix[j][colStart])
			}
		}
		rowStart++
		colStart++
		rowEnd--
		colEnd--
	}
	return ans
}
