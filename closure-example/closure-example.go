package main

import (
	"fmt"
)


func main() {
	
	next := MiFuncion()
	for i:=0; i<7; i++ {
		fmt.Printf("client fetched '%v'\n", next())
	}
	

}

func MiFuncion() func() string {

    i := 0
    var data []string = []string {"shura", "ikki", "shiryu", "seiya", "shun", "aldebaran", "nachi"}

    return func()string{
       val := data[i]
	i++
	return val
    }
}

