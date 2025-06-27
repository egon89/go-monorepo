package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("> ex2")

	ch := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, time.Duration(i*100)*time.Millisecond, &wg, ch)
	}

	go closer(&wg, ch)

	for result := range ch {
		fmt.Printf("result: %s\n", result)
	}

	fmt.Println("# end")
}

func worker(id int, duration time.Duration, wg *sync.WaitGroup, ch chan string) {
	// wg.Add(1) // it doesn't work here
	defer wg.Done()

	fmt.Printf("worker %d running...\n", id)

	time.Sleep(duration)

	ch <- fmt.Sprintf("message id %d", id)
}

func closer(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("closer running...")

	wg.Wait()

	close(ch)
}

/*
	> ex2
	closer running...
	worker 0 running...
	result: message id 0
	worker 6 running...
	worker 7 running...
	worker 8 running...
	worker 9 running...
	worker 2 running...
	worker 5 running...
	worker 1 running...
	worker 3 running...
	worker 4 running...
	result: message id 1
	result: message id 2
	result: message id 3
	result: message id 4
	result: message id 5
	result: message id 6
	result: message id 7
	result: message id 8
	result: message id 9
	# end
*/
