package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("> ex1")
	ch := make(chan string)

	go worker(1, 2*time.Second, ch)

	result := <-ch
	fmt.Printf("result: %s\n", result)
}

func worker(id int, duration time.Duration, ch chan string) {
	time.Sleep(duration)
	ch <- fmt.Sprintf("send message from worker %d\n", id)
}
