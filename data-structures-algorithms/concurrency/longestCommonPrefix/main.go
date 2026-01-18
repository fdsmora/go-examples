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
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	chChan := make(chan element)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Buffered channel to prevent blocking when workers need to acknowledge
	wait := make(chan struct{}, len(strs))

	var wg sync.WaitGroup
	for _, w := range strs {
		wg.Add(1)
		go worker(w, wait, ctx, cancel, chChan, &wg)
	}

	// Close channels when done to signal workers
	go func() {
		wg.Wait()
		close(chChan)
		close(wait)
	}()

	// Check each position in the first string
	for i, ch := range strs[0] {
		// Send the character to all workers
		for range len(strs) {
			select {
			case <-ctx.Done():
				return strs[0][0:i]
			case chChan <- element{ch, i}:
				// Sent successfully
			}
		}

		// Wait for all workers to acknowledge
		for range len(strs) {
			select {
			case <-ctx.Done():
				return strs[0][0:i]
			case <-wait:
				// Worker acknowledged
			}
		}
		fmt.Printf("Position %d ('%c'): All workers matched\n", i, ch)
	}

	// All characters matched - return the entire first string
	return strs[0]
}

func worker(word string, wait chan<- struct{}, ctx context.Context, cancel context.CancelFunc, chChan <-chan element, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, c := range word {
		select {
		case elem, ok := <-chChan:
			if !ok {
				// Channel closed - we're done
				return
			}

			// Check if character matches
			if i != elem.i || c != elem.ch {
				fmt.Printf("Worker '%s': Mismatch at position %d (expected '%c', got '%c')\n",
					word, i, c, elem.ch)
				cancel()
				return
			}

			fmt.Printf("Worker '%s': Match at position %d ('%c')\n", word, i, elem.ch)

			// Acknowledge match - use select to avoid blocking if context is cancelled
			select {
			case wait <- struct{}{}:
				// Acknowledged successfully
			case <-ctx.Done():
				return
			}

		case <-ctx.Done():
			return
		}
	}

	// This word is shorter - cancel since we can't match further
	fmt.Printf("Worker '%s': Reached end (length %d)\n", word, len(word))
	cancel()
}
