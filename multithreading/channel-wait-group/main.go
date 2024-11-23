package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channel with wait group")

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(5)

	go publish(ch)
	go read(ch, &wg)

	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func read(ch chan int, wg *sync.WaitGroup) {
	for value := range ch {
		fmt.Printf("read: %v\n", value)
		wg.Done()
	}
}

/*
channel with wait group
read: 0
read: 1
read: 2
read: 3
read: 4
*/
