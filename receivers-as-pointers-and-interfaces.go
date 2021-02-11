/*
https://play.golang.org/p/UDPbLxndkDj

Pareciera independiente de la interfaz, que la regla para la funcion es
1. Si el receiver no es pointer le puedes pasar pointer o copia (por valor)
2. Si el receiver es pointer, solo le puedes pasar pointer, si pasas por valor falla
*/
package main

import (
	"fmt"
)

type TestInterface interface{
	Func1() string
}

type Struct struct{}
func (s Struct) Func1() string {
	return "Hello World"
}

type StructPtr struct{}
func (s *StructPtr) Func1() string {
	return "Hello World"
}

func main() {
	var x TestInterface
	
	// works
	x = Struct{}
	fmt.Println(x.Func1())
	
	// works
	x = &Struct{}
	fmt.Println(x.Func1())
	
	// fails: StructPtr does not implement TestInterface (Func1 method has pointer receiver)
	// x = StructPtr{}
	// fmt.Println(x.Func1())
	
	// works
	x = &StructPtr{}
	fmt.Println(x.Func1())
}
