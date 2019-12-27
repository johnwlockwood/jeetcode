package rotate

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < n; i++ {
		low := 0
		high := n - 1
		for l, r := low, high; l < r; l++ {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			r--
		}
	}
}
