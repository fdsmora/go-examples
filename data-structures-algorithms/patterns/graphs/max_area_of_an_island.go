package graphs

// leetcode problem 695
// My solution, time complexity is O(n*m) where n and m are the dimensions of the input grid, as we traverse each cell of the grid once. Space complexity is O(n*m) in the worst case when the stack contains all cells of the grid.
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	st := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = -1
				st = append(st, []int{i, j})
				area := 0
				for len(st) > 0 {
					x := st[len(st)-1][0]
					y := st[len(st)-1][1]
					st = st[:len(st)-1]
					area++
					explore(grid, x, y, &st)
				}
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}

func explore(grid [][]int, i, j int, st *[][]int) {
	if i-1 >= 0 && grid[i-1][j] == 1 {
		grid[i-1][j] = -1
		*st = append(*st, []int{i - 1, j})
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
		grid[i][j+1] = -1
		*st = append(*st, []int{i, j + 1})
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		grid[i+1][j] = -1
		*st = append(*st, []int{i + 1, j})
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		grid[i][j-1] = -1
		*st = append(*st, []int{i, j - 1})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
