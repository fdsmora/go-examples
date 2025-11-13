package goconcurrencycourse

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch, quit := make(chan string), make(chan string)

	wg.Add(1)
	for i := 0; i < 3; i++ {
		go Race(ch, quit, i)
	}

	for {
		select {
		case update := <-ch:
			fmt.Println(update)
		case finish := <-quit:
			fmt.Println(finish)
			quit <- "you win!"
			wg.Wait()
			return
		}
	}
}

func Race(ch, quit chan string, i int) {
	ch <- fmt.Sprintf("car %d started!", i)
	for {
		time.Sleep(time.Duration(rand.IntN(3)) * time.Millisecond)
		ch <- fmt.Sprintf("car %d running", i)
		quit <- fmt.Sprintf("car %d finished", i)
		fmt.Println(<-quit)
		wg.Done()
	}

}
