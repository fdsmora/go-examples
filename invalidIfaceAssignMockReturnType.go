package main

import (
	"fmt"
)

type (
	I interface {
		f() (R, error)
	}
	A struct{}

	R interface {
		r() int
	}
	B struct{}
)

func (a A) f() (B, error) {
	return B{}, nil
}
func (b B) r() int {
	return 1
}

// Will throw error because even the returned B from f() implements R, for some reason I don't understand it is not allowed
func main() {
	var a I = A{}
	b, _ := a.f()
	fmt.Println(b.r())
}

