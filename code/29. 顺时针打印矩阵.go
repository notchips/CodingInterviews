func printMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)
	travelEdge(matrix, &ans)
	return ans
}

func travelEdge(matrix [][]int, ans *[]int) {
	m, n := len(matrix), len(matrix[0])
	for col := 0; col < n; col++ {
		*ans = append(*ans, matrix[0][col])
	}
	for row := 1; row < m-1; row++ {
		*ans = append(*ans, matrix[row][n-1])
	}
	if m > 1 {
		for col := n - 1; col >= 0; col-- {
			*ans = append(*ans, matrix[m-1][col])
		}
	}
	if n > 1 {
		for row := m - 2; row >= 1; row-- {
			*ans = append(*ans, matrix[row][0])
		}
	}
	if m > 2 && n > 2 {
		matrix = matrix[1 : m-1]
		for row := range matrix {
			matrix[row] = matrix[row][1 : n-1]
		}
		travelEdge(matrix, ans)
	}
}
