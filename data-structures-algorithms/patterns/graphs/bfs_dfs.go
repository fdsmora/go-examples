package graphs

import (
	"github.com/fdsmora/go-examples/ds-algorithms/queue"
	"github.com/fdsmora/go-examples/ds-algorithms/stack"
)

/*
type Queue[T any] []T

func (q *Queue[T]) Add(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		var zero T
		return zero
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) Length() int {
	return len(*q)
}
*/

type Node struct {
	Value    int
	Adjacent []*Node
	Visited  bool
}

func BFS(r *Node) {
	q := queue.Queue[*Node]{}

	visited := make(map[*Node]bool)
	q.Enqueue(r)
	visited[r] = true

	for len(q) > 0 {
		c := q.Dequeue()
		visited[c] = true
		// do something with c
		for _, n := range c.Adjacent {
			if !visited[c] {
				visited[c] = true
				q.Enqueue(n)
			}
		}
	}
}

func DFSIterative(r *Node) {
	stack := stack.Stack[*Node]{}
	visited := make(map[*Node]bool)
	stack.Push(r)
	visited[r] = true
	for !stack.IsEmpty() {
		c, _ := stack.Pop()
		visited[c] = true
		// process c
		for _, n := range c.Adjacent {
			if !visited[n] {
				visited[n] = true
				stack.Push(n)
			}
		}
	}
}

func DFS(r *Node) {
	visited := map[*Node]bool{}
	visited[r] = true
	DFS_(r, visited)
}

func DFS_(r *Node, visited map[*Node]bool) {
	if r.Visited {
		return
	}
	r.Visited = true
	for _, n := range r.Adjacent {
		// process n
		if !visited[n] {
			DFS_(n, visited)
		}
	}
}

/*
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}*/
