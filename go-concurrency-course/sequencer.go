package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Cook struct {
	Name string
	Wait chan bool
}

func main() {

	fCh := fanIn(cookFood("Player 1: "), cookFood("Player 2: "))

	for i := 0; i < 3; i++ {
		player1 := <-fCh
		fmt.Println(player1.Name)

		player2 := <-fCh
		fmt.Println(player2.Name)

		player1.Wait <- true
		player2.Wait <- true

		fmt.Printf("Done with %d round\n", i)
	}
}

func cookFood(name string) <-chan Cook {
	wait := make(chan bool)
	outCh := make(chan Cook)
	go func() {
		for i := 0; ; i++ {
			outCh <- Cook{fmt.Sprintf("%s %s", name, "Done"), wait}
			time.Sleep(time.Duration(rand.IntN(3)) * time.Millisecond)
			<-wait
		}
	}()
	return outCh
}

func fanIn(ch1, ch2 <-chan Cook) <-chan Cook {
	outCh := make(chan Cook)

	go func() {
		for {
			select {
			case cook1 := <-ch1:
				outCh <- cook1
			case cook2 := <-ch2:
				outCh <- cook2
			}
		}
	}()

	return outCh
}
