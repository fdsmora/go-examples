package chapter1

func RotateMatrix(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 || len(matrix[0]) != n {
		return false
	}
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j] // save top
			// left -> top
			matrix[i][j] = matrix[n-j-1][i]
			// bottom -> left
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			// right -> bottom
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			// top -> right
			matrix[j][n-i-1] = temp
		}
	}
	return true
}
