func hasPath(matrix [][]byte, str string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for row := 0; row < len(matrix); row ++ {
		for col := 0; col < len(matrix[0]); col++ {
			if search(matrix, row, col, str, 0) {
				return true
			}
		}
	}
	return false
}

func search(matrix [][]byte, row, col int, str string, i int) bool {
	if i == len(str) {
		return true
	}
	if row < 0 || row >= len(matrix) ||
		col < 0 || col >= len(matrix[0]) ||
		matrix[row][col] != str[i] {
		return false
	}
	matrix[row][col] ^= 1 << 7
	ret := search(matrix, row-1, col, str, i+1) ||
		search(matrix, row+1, col, str, i+1) ||
		search(matrix, row, col-1, str, i+1) ||
		search(matrix, row, col+1, str, i+1)
	matrix[row][col] ^= 1 << 7
	return ret
}