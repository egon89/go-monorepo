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

	// even numbers from 0 to 10
	for i := 0; i <= 10; i += 2 {
		wg.Add(1)
		go worker(i, "even", time.Duration(i*100)*time.Millisecond, &wg, ch)
	}

	// odd numbers from 1 to 9
	for i := 1; i < 10; i += 2 {
		wg.Add(1)
		go worker(i, "odd", time.Duration(i*100)*time.Millisecond, &wg, ch)
	}

	go closer(&wg, ch)

	for result := range ch {
		fmt.Printf("result: %s\n", result)
	}

	fmt.Println("# end")
}

func worker(id int, prefix string, duration time.Duration, wg *sync.WaitGroup, ch chan string) {
	// wg.Add(1) // it doesn't work here. It should be called outside of the goroutine
	defer wg.Done()

	fmt.Printf("worker %d running...waiting %v ms\n", id, duration.Milliseconds())

	time.Sleep(duration)

	ch <- fmt.Sprintf("%s: message id %d", prefix, id)
}

func closer(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("closer running...")

	wg.Wait()

	close(ch)
}

/*
> ex2
	closer running...
	worker 0 running...waiting 0 ms
	result: even: message id 0
	worker 1 running...waiting 100 ms
	worker 10 running...waiting 1000 ms
	worker 5 running...waiting 500 ms
	worker 7 running...waiting 700 ms
	worker 9 running...waiting 900 ms
	worker 4 running...waiting 400 ms
	worker 2 running...waiting 200 ms
	worker 6 running...waiting 600 ms
	worker 3 running...waiting 300 ms
	worker 8 running...waiting 800 ms
	result: odd: message id 1
	result: even: message id 2
	result: odd: message id 3
	result: even: message id 4
	result: odd: message id 5
	result: even: message id 6
	result: odd: message id 7
	result: even: message id 8
	result: odd: message id 9
	result: even: message id 10
	# end
*/
