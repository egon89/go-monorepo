package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64

func main() {
	fmt.Println("load balancer")
	requests := 1000000
	workers := 2
	wg := sync.WaitGroup{}
	wg.Add(requests)
	ch := make(chan int)

	for i := 1; i <= workers; i++ {
		go worker(i, ch, &wg)
	}

	startTime := time.Now()
	dispatcher(requests, ch)
	wg.Wait()

	fmt.Printf("counter: %v\n", counter)
	fmt.Printf("time: %v\n", time.Since(startTime))
}

func dispatcher(requests int, ch chan<- int) {
	for i := 0; i < requests; i++ {
		ch <- i
	}
}

func worker(workerId int, ch <-chan int, wg *sync.WaitGroup) {
	for value := range ch {
		fmt.Printf("worker %v: value: %v\n", workerId, value)
		atomic.AddInt64(&counter, 1)
		wg.Done()
	}
}
