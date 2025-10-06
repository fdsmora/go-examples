package chapter1

func SetToZeros(mtx [][]int) {
	m := len(mtx)
	n := len(mtx[0])

	var rowHasZeros, colHasZeros bool

	for c := 0; c < n; c++ {
		if mtx[0][c] == 0 {
			rowHasZeros = true
			break
		}
	}
	for r := 0; r < m; r++ {
		if mtx[r][0] == 0 {
			colHasZeros = true
			break
		}
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if mtx[r][c] == 0 {
				mtx[0][c] = 0
				mtx[r][0] = 0
			}
		}
	}

	for r := 1; r < m; r++ {
		if mtx[r][0] == 0 {
			nullifyRow(mtx, r)
		}
	}
	for c := 1; c < n; c++ {
		if mtx[0][c] == 0 {
			nullifyColumn(mtx, c)
		}
	}
	if rowHasZeros {
		nullifyRow(mtx, 0)
	}
	if colHasZeros {
		nullifyColumn(mtx, 0)
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
	for c := 0; c < len(mtx[0]); c++ {
		mtx[row][c] = 0
	}
}
