package chapter1

func SetToZeros(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var zeroFound bool

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if matrix[r][c] == 0 {
				zeroFound = true
				matrix[0][c] = -1
				matrix[r][0] = -1
			}
		}
	}
	if !zeroFound {
		return
	}
	for r := 0; r < m; r++ {
		if matrix[r][0] == -1 {
			for c := 0; c < n; c++ {
				if r == 0 && matrix[r][c] == -1 {
					continue
				}
				matrix[r][c] = 0
			}
		}
	}
	for c := 0; c < n; c++ {
		if matrix[0][c] == -1 {
			for r := 0; r < m; r++ {
				matrix[r][c] = 0
			}
		}
	}
}

// Limitation is that works only with m*n matrices where m and n <=32, so space is O(1).
// If n and m are unbounded, the space complexity is O(m+n) space.
func SetToZerosBitVector(mtx [][]int) {
	m, n := len(mtx), len(mtx[0])

	var rows, columns int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mtx[r][c] == 0 {
				rows |= 1 << r
				columns |= 1 << c
			}
		}
	}
	for r := 0; r < m; r++ {
		if rows&(1<<r) > 0 {
			nullifyRow(mtx, r)
		}
	}
	for c := 0; c < n; c++ {
		if columns&(1<<c) > 0 {
			nullifyColumn(mtx, c)
		}
	}
}

func nullifyColumn(mtx [][]int, col int) {
	for r := 0; r < len(mtx); r++ {
		mtx[r][col] = 0
	}
}

func nullifyRow(mtx [][]int, row int) {
	for c := 0; c < len(mtx); c++ {
		mtx[row][c] = 0
	}
}
