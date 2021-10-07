package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	c chan string
        data []string = []string {"shura", "ikki", "shiryu", "seiya", "shun", "aldebaran", "nachi"}
	num int
)

func main() {
	
	for i:=0; i<len(data); i++ {
		fmt.Printf("client fetched '%v'\n", MiFuncion())
	}
	

}

func MiFuncion()string {

    once.Do(func(){
	fmt.Println("inside once")	
        c = make(chan string)
	go func() {
		fmt.Printf("go rutine, data: %v, chan: %v\nNotice how the goroutine has access to the global variables ^(really?)\n", data, c)
		for _, item := range data {
	    		c <- item
		    fmt.Printf("inner loop '%s'\n", item)
		}
		close(c)
		
	}()
    })
    res := <- c
    return res
}

