package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8}
	b := []float32{1,2,3,4,5,6,7,8}
	
	// invalid operation: a_p[3] (type *[]int does not support indexing)
	//a_p := &a
	//val := a_p[3]
	
	fmt.Printf("%v\n", reflect.DeepEqual(a,b) )
}

