package problem4_1

// Problem 4.1

type Node struct {
	Value      int
	Neighbours []*Node
}

func FindRoute(s, e *Node) bool {
	visited := map[*Node]bool{}
	if s == e {
		return true
	}
	return DFS(s, e, visited)
}

func DFS(c, e *Node, visited map[*Node]bool) bool {
	visited[c] = true
	for _, n := range c.Neighbours {
		if _, ok := visited[n]; ok {
			continue
		}
		if n == e || DFS(n, e, visited) {
			return true
		}
	}
	return false
}
