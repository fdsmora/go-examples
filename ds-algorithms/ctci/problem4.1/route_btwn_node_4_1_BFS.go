package problem4_1

import "github.com/fdsmora/go-examples/ds-algorithms/queue"

type Status int

const (
	Unvisited Status = iota
	Visited
	Visiting
)

type NodeB struct {
	Value      int
	Status     Status
	Neighbours []*NodeB
}

func FindRouteBFS(s, e *NodeB) bool {
	Q := &queue.Queue[*NodeB]{}
	Q.Enqueue(s)
	s.Status = Visiting // Mark the start node as visiting

	for len(*Q) > 0 {
		c := Q.Dequeue()
		c.Status = Visited // Mark the current node as visited

		// If we found the destination node, return true
		if e == c {
			return true
		}

		// Traverse all neighbors of the current node
		for _, n := range c.Neighbours {
			if n.Status == Unvisited {
				n.Status = Visiting // Mark as visiting
				Q.Enqueue(n)
			}
		}
	}

	// No path found
	return false
}
