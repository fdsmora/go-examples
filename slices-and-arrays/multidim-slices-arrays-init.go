package main

// Multiarray & multislice initialization
// See https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
// http://www.golangbootcamp.com/book/collection_types

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	var slc = make([][]int, 5) // slices are descriptors
	var arr [3][3]int //arrays are values
	fmt.Println(slc)
	fmt.Println(arr)
	fmt.Printf("slc is type %T\n", slc)
	fmt.Printf("slc is type %T\n", arr)
}
