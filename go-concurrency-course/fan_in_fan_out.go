package goconcurrencycourse

package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	var nums [10]int
	for i := 0; i < 10; i++ {
		nums[i] = rand.IntN(500)
	}

	ch := generator1(nums)

	double1 := double(ch)
	double2 := double(ch)

	fCh := fanIn(double1, double2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-fCh)
	}
}

func double(ch <-chan string) chan string {

	outCh := make(chan string)

	go func() {
		for v := range ch {
			k, _ := strconv.Atoi(v)
			outCh <- fmt.Sprintf("%s * 2 = %d", v, k*2)
		}
		close(outCh)
	}()

	return outCh
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-ch1:
				ch <- msg1
			case msg2 := <-ch2:
				ch <- msg2
			}
		}
	}()

	return ch
}

func generator1(nums [10]int) chan string {
	ch := make(chan string)

	go func() {
		for v := range nums {
			ch <- strconv.Itoa(v)
		}
		close(ch)
	}()

	return ch
}
