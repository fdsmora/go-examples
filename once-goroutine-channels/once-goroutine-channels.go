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

/* Another equivalent approach without `once` */

/*
func TestMiFuncion(t *testing.T) {
	fmt.Println("Hello, playground")

	//for i := 0; i < 5; i++ {
	//	fmt.Println(MiFuncion())
	//}

	for v := range MiFuncion() {
		fmt.Printf("valor sacado: %v\n", v)
	}

}

func MiFuncion() chan string {

	var c chan string
	var data []string = []string{"shura", "ikki", "shiryu", "seiya", "shun", "aldebaran", "nachi"}
	//	once.Do(func() {
	//		fmt.Println("once")
	c = make(chan string)
	go func() {
		fmt.Printf("go rutine, data: %v, chan: %v\n", data, c)
		for _, item := range data {
			fmt.Printf("inner loop %s\n", item)
			//time.Sleep(10 * time.Second)
			c <- item
		}
		close(c)
		// 			for {
		//	time.Sleep(3 * time.Second)
//			fmt.Println("goroutine inside once alive...")
	//	} 

	}()
	//	})
	//time.Sleep(13 * time.Second)
	//res := "hola" + strconv.Itoa(num)
	//res := <-c
	//fmt.Println("fetched from channel: " + res)
	return c
}
*/
