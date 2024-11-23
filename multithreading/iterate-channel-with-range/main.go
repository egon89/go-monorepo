package main

import "fmt"

func main() {
	fmt.Println("channel")
	ch := make(chan int)

	go func() {
		fmt.Println("goroutine working")
		ch <- 1
		close(ch) // we need to close the channel to avoid deadlock
	}()

	for value := range ch {
		fmt.Printf("read: %v\n", value)
	}
}

/*
goroutine working
read: 1
*/
