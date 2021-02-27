package main

// https://stackoverflow.com/questions/25543520/declare-slice-or-make-slice

import (
	"fmt"
)

func main() {
	// zero slices
	var slc []int // The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.
 	fmt.Println("Zero value of slc:", slc) // The zero value of a slice (nil) acts like a zero-length slice
	fmt.Printf("slc type: %T\n", slc)
	fmt.Println("slc is nil? ", slc==nil) 
	fmt.Printf("len of slc: %d\ncapacity of slc: %d\n", len(slc), cap(slc))
	
	
	fmt.Println(Filter([]int{1,2,3,4,5,6,7,8}, func (i int) bool { return i % 2 == 1 }))
}

// Since the zero value of a slice (nil) acts like a zero-length slice, you can declare a slice variable and then append to it in a loop:
// Additional note on `s`: here we pass the `s` slice by value, because slices are simple, cheap types that are just some kind of "header" or roughly "metadata" to an array,
// An slice just has a pointer to an underlying array, a lenght and a capacity. So it doesn't matter if all time slices are passed by value, its always cheap. 
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
