package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("> ex3")
	var wg sync.WaitGroup

	// semaphores for each pair: even->odd
	pairs := 5 // 0-1, 2-3, 4-5, 6-7, 8-9
	semaphores := make([]chan struct{}, pairs)

	for i := 0; i < pairs; i++ {
		semaphores[i] = make(chan struct{})
	}

	for i := 0; i < pairs; i++ {
		even := i * 2
		odd := even + 1

		wg.Add(1)
		go func(value int, channel chan struct{}) {
			defer wg.Done()

			fmt.Printf("# working on even value %d\n", value)

			sleep(value)

			fmt.Printf("## even value completed: %d\n", value)

			channel <- struct{}{} // release odd

			// we can use close(channel) too
		}(even, semaphores[i])

		wg.Add(1)
		go func(value int, channel chan struct{}) {
			defer wg.Done()

			fmt.Printf("# waiting for even value %d\n", even)

			<-channel // wait for even to complete

			fmt.Printf("## working on odd value %d [pair %d-%d] \n", value, even, value)
		}(odd, semaphores[i])

	}

	wg.Wait()

	fmt.Println("...done")
}

func sleep(value int) {
	duration := time.Duration(value) * time.Millisecond

	fmt.Printf("...value %d will sleep %v ms\n", value, duration)

	time.Sleep(duration)
}
