package main

type Node struct {
	Value int
}

type Queue []*Node

func (q *Queue) Dequeue() *Node {
	if len(*q) == 0 {
		return nil
	}
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func (q *Queue) Enqueue(c *Node) {
	*q = append(*q, c)
}
