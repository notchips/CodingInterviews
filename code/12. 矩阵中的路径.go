package code

func HasPath(matrix [][]byte, str string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if search(matrix, str, row, col) {
				return true
			}
		}
	}
	return false
}

func search(matrix [][]byte, str string, row, col int) bool {
	if len(str) == 0 {
		return true
	}
	if row < 0 || row >= len(matrix) ||
		col < 0 || col >= len(matrix[0]) ||
		matrix[row][col] != str[0] {
		return false
	}
	matrix[row][col] ^= 1 << 7
	ret := search(matrix, str[1:], row+1, col) ||
		search(matrix, str[1:], row-1, col) ||
		search(matrix, str[1:], row, col+1) ||
		search(matrix, str[1:], row, col+1)
	matrix[row][col] ^= 1 << 7
	return ret
}
