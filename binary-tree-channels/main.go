package main

import (
	"binary-tree-channels/tree"
	"fmt"
)

func main() {
	t := tree.New(1)

	//	t.Print()
	c := make(chan int)
	go t.Walk(c)

	for v := range c {
		fmt.Printf(" %d", v)
	}
}
