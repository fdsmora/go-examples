package main

import (
	"fmt"
	"encoding/json"
)

var data string = `{ "Field1" : { "Attrib" : "abcde" } }`

type X struct {
    Field1 A 
    Field2 A
}
 
type A struct {
    Attrib string
}

func main() {
	
	// Attempt to "nullify" an struct's field
	// It's not possible since the field is not declared as pointer
	// So it will take blank (zero) value of struct type when not assigned
	/*x := X{
		Field1: A{ "hola"},
	 }
	
	fmt.Printf("%v\n", x)
	
	ref := &x.Field1
	
	fmt.Printf("Ref: %v\n", ref)
	
	ref = nil
	fmt.Printf("%v\n", x)*/
	
	
	
	// Also, when parsing a JSON string into an struct, the missing fields will be left as zero values of struct
	var x X
	
	json.Unmarshal([]byte(data), &x)
	fmt.Printf("%v\n", x)
}

