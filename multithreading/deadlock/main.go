package main

import "fmt"

func main() {
	fmt.Println("deadlock example")
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("i:%v\n", i)
		}
	}()

	// no goroutines to fill the channel
	<-ch
}

/*
deadlock example
i:0
i:1
i:2
i:3
i:4
fatal error: all goroutines are asleep - deadlock
*/
