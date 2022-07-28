package main

// Example of capturing variables caveant inside closure:
// https://notes.shichao.io/gopl/ch5/#caveat-capturing-iteration-variables

import (
	"fmt"
)

type Cosa struct {
	cont string
}

func main() {

	var values []*Cosa = []*Cosa{
		&Cosa{"A"},
		&Cosa{"B"},
		&Cosa{"C"},
	}

	fmt.Printf("APUNTADOR: %p \n", values[0])
	fmt.Printf("APUNTADOR: %p \n", values[1])
	fmt.Printf("APUNTADOR: %p \n", values[2])

	var anyFuncions []func()

	for _, d := range values {
		fmt.Printf("d value is %p\n", &d)
		dir := d // NOTE: necessary to avoid subtle issue!
		fmt.Printf("dir value is %p\n", &dir)
		//	dir := d // NOTE: necessary!
		anyFuncions = append(anyFuncions, func() {
			Operating(dir)
		})
	}

	for _, anyFunc := range anyFuncions {
		anyFunc() // clean up
	}
}

func Operating(c *Cosa) {
	fmt.Printf("c %p\n", &c)
	//fmt.Printf("operating: %#v\n", c) // Do anything with c
}

