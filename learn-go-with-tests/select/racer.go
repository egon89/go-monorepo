package main

import (
	"net/http"
	"time"
)

// named return value
// if you don't assign anything to winner, it will return the zero value for a string, which is ""
func Racer(a, b string) (winner string) {
	return
}

// V1
func RacerV1(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}

// V2
// select allows to wait on multiple channels
// the first one to send a value wins and the code underneath the case is executed
func RacerV2(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// struct{} is the smallest data type available from a memory perspective
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
