package main

import (
	"binary-tree-channels/tree"
	"fmt"
)

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go t1.Walk(c1)
	go t2.Walk(c2)

	for {
		x1, ok1 := <-c1
		x2, ok2 := <-c2
		switch {
		case ok1 != ok2:
			return false
		case !ok1:
			return true
		case x1 != x2:
			return false
		default:
		}
	}
}

func main() {
	t := tree.New(1)

	//	t.Print()
	c := make(chan int)
	go t.Walk(c)

	for v := range c {
		fmt.Printf(" %d", v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
}
