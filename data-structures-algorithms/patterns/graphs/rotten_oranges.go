package graphs

// leetcode problem 695

// 1st attempt, inneficient, using circular linked list and round robin on the rotten oranges
/*

type (
	Coords struct {
		X,Y, dist int
	}
	RottenOrange struct {
		curr *Coords
		Q []*Coords
		next, prev *RottenOrange
	}
	List struct{

	}
)

const (
	zero int = iota
	fresh
    rotten
)

func orangesRotting(grid [][]int) int {
    var maxDist = -1
	var head, areFruits = getListOfRottenOranges(grid)
	if !areFruits {
		return 0
	}
	if head == nil {
		return -1
	}
	var cOrange = head
	// debug


	for cOrange != nil {
		// debug
		//fmt.Printf("queue size: %d\n", len(cOrange.Q))
		if len(cOrange.Q) > 0 {
			curr := cOrange.Q[0]
			cOrange.Q = cOrange.Q[1:]
			addToQueue(grid, curr, &cOrange.Q)
			maxDist = max(maxDist, curr.dist)
		}else{
			if cOrange.next == cOrange {
				cOrange.next = nil
			}else{
				cOrange.prev = cOrange.next
				cOrange.next.prev = cOrange.prev
			}
		}
		cOrange = cOrange.next
	}
	return maxDist
}

func max (a,b int)int{
	if a > b {
		return a
	}
	return b
}

func addToQueue(grid [][]int, curr *Coords, Q *[]*Coords) {
	var R, C = len(grid), len(grid[0])
	var x, y, dist = curr.X, curr.Y, curr.dist
	// debug
	//fmt.Println("fresh:",fresh)
	//fmt.Printf("coords: %d\n", grid[x][y])
	//fmt.Printf("grid[x][y+1]: %d\n", grid[x][y+1])
	if y+1<C && grid[x][y+1] == fresh {
		*Q = append(*Q, &Coords{x, y+1, dist+1})
		grid[x][y+1]=rotten
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
	//fmt.Printf("grid[x+1][y]: %d\n", grid[x+1][y])
	if x+1<R && grid[x+1][y] == fresh{
		*Q = append(*Q, &Coords{x+1,y,dist+1})
		grid[x+1][y]=rotten
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
	//	fmt.Printf("grid[x][y-1]: %d\n", grid[x][y-1])
	if y-1>=0 && grid[x][y-1] == fresh{
		*Q = append(*Q, &Coords{x,y-1,dist+1})
		grid[x][y-1]=rotten
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
//		fmt.Printf("grid[x-1][y]: %d\n", grid[x-1][y])

	if x-1>=0 && grid[x-1][y] == fresh{
		*Q = append(*Q, &Coords{x-1,y,dist+1})
		grid[x-1][y]=rotten
	// debug
	//fmt.Printf("successfully added to queue: %#v\n", *Q)
	}
//	fmt.Printf("successfully added to queue: %#v\n", *Q)

	//return dist + 1
}

func getListOfRottenOranges(grid [][]int) (*RottenOrange, bool) {
	var R, C = len(grid), len(grid[0])
	var head, curr *RottenOrange
	var areFruits bool
	// debug
	//fmt.Println("About to create list")
	// visit every orange in the grid in order to find rotten oranges
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++{
			if grid[i][j]==rotten {
				areFruits = true
				// 1st rotten orange found case
				//fmt.Printf("rotten found at (%d,%d)\n", i, j)
				if curr == nil {
					newCoords := &Coords{i,j,0}
					head = &RottenOrange{
						curr: newCoords,
						Q : []*Coords{newCoords},
					}
					curr = head
				//	fmt.Printf("head is %v\n", curr)
				// not 1st rotten orange found
				}else{
					newCoords := &Coords{i,j,0}
					curr.next = &RottenOrange{
						curr: newCoords,
						Q: []*Coords{newCoords},
						prev: curr,
					}
					curr = curr.next
				}
			}else if grid[i][j]==fresh {
				areFruits = true
				// unreachable fruit
				if ((j+1 < C && grid[i][j+1]==zero) || j+1==C) && // right
					((i+1 < R && grid[i+1][j]==zero) || i+1==R) && // down
					((j-1>=0 && grid[i][j-1]==zero) || j-1<0) && //left
					((i-1>=0 && grid[i-1][j]==zero) || i-1<0) {
					return nil, areFruits
				}
			}
		}
	}

	if head != nil {
		// make list circular
		curr.next = head
		head.prev = curr
	}

	return head, areFruits
}

*/

// my correct solution, time complexity is O(n*m) where n and m are the dimensions of the input grid, as we traverse each cell of the grid once. Space complexity is O(n*m) in the worst case when the stack contains all cells of the grid.
type (
	Coords struct {
		X, Y, dist int
	}
)

const (
	zero int = iota
	fresh
	rotten
)

func orangesRotting(grid [][]int) int {
	var maxDist = 0

	Q := []Coords{}
	var R, C = len(grid), len(grid[0])
	freshCount := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == rotten {
				Q = append(Q, Coords{i, j, 0})
			} else if grid[i][j] == fresh {
				freshCount++
			}
		}
	}

	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]
		maxDist = max(maxDist, curr.dist)
		freshCount = rot(grid, &curr, &Q, freshCount)
	}

	if freshCount > 0 {
		return -1
	}

	return maxDist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rot(grid [][]int, curr *Coords, Q *[]Coords, freshCount int) int {
	var R, C = len(grid), len(grid[0])
	var x, y, dist = curr.X, curr.Y, curr.dist

	if y+1 < C && grid[x][y+1] == fresh {
		*Q = append(*Q, Coords{x, y + 1, dist + 1})
		grid[x][y+1] = rotten
		freshCount--
	}

	if x+1 < R && grid[x+1][y] == fresh {
		*Q = append(*Q, Coords{x + 1, y, dist + 1})
		grid[x+1][y] = rotten
		freshCount--
	}
	if y-1 >= 0 && grid[x][y-1] == fresh {
		*Q = append(*Q, Coords{x, y - 1, dist + 1})
		grid[x][y-1] = rotten
		freshCount--
	}

	if x-1 >= 0 && grid[x-1][y] == fresh {
		*Q = append(*Q, Coords{x - 1, y, dist + 1})
		grid[x-1][y] = rotten
		freshCount--

	}
	return freshCount
}
