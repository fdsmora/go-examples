package longestcommonprefix

import (
	"context"
	"fmt"
	"sync"
)

type element struct {
	ch rune
	i  int
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	chChan := make(chan element)
	//	abort := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wait := make(chan struct{})
	defer close(wait)
	var wg sync.WaitGroup
	for _, w := range strs {
		wg.Add(1)
		go worker(w, wait, ctx, cancel, chChan, &wg)
	}
	for i, ch := range strs[0] {
		for range n {
			select {
			case <-ctx.Done():
				return strs[0][0:i]
			default:
				chChan <- element{ch, i}
			}
		}
		for range len(strs) {
			select {
			case <-ctx.Done():
				return strs[0][0:i]
			case <-wait:
				continue
			}
		}
		fmt.Println("-----------------------")
	}
	wg.Wait()
	fmt.Println("All workers completed successfully")
	return ""
}

func worker(word string, wait chan<- struct{}, ctx context.Context, cancel context.CancelFunc, chChan <-chan element, wg *sync.WaitGroup) {
	defer wg.Done()
	var elem element
	for i, c := range word {
		select {
		case elem = <-chChan:
			if i != elem.i || c != elem.ch {
				cancel()
				return
			}
		case <-ctx.Done():
			return
		}
		fmt.Printf("Goroutine '%s' got: '%s'\n", word, string(elem.ch))
		wait <- struct{}{}
	}
}
