package main

import (
	"context"
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=uiUCIz-3CWM

// use of context
// controlling timeouts
// cancel goroutines
// passing metadata across your application (middleware, authorizations, etc)

func main() {
	timeoutExample()
	timeoutWithStringExample()
	passInformationAcrossSystemWithContext()
	passInformationAcrossSystemWithContextUsingStruct()
}

func timeoutExample() {
	fmt.Println("> timeoutExample")
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	// struct{} represents an empty struct (no space in memory)
	// we cannot add fields to an empty struct{}
	done := make(chan struct{})

	// simulate a third party call
	go func() {
		time.Sleep(400 * time.Millisecond) // prints: api called
		// time.Sleep(401 * time.Millisecond) // prints: timeout error
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("api called")
	case <-ctx.Done():
		fmt.Println("timeout error")
	}
}

func timeoutWithStringExample() {
	fmt.Println("> timeoutWithStringExample")
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	stringChannel := make(chan string)

	go func() {
		time.Sleep(700 * time.Millisecond)
		stringChannel <- "john doe"
		close(stringChannel)
	}()

	select {
	case value, hasValue := <-stringChannel:
		if hasValue {
			fmt.Printf("value: %v\n", value)
		} else {
			fmt.Println("string channel closed")
		}
	case <-ctx.Done():
		fmt.Println("timeout error")
	}

	// value: john doe
}

// Use context values only for request-scoped data that transits
// processes and API boundaries, not for passing optional parameters to
// functions.
func passInformationAcrossSystemWithContext() {
	fmt.Println("> passInformationAcrossSystemWithContext")
	const userId = "user-id"
	ctx := context.WithValue(context.Background(), userId, "uuid")
	if id, ok := ctx.Value(userId).(string); ok {
		fmt.Printf("the value is %v\n", id)
	} else {
		fmt.Println("no value in context")
	}

	// the value is uuid
}

type RequestContext struct {
	UserId  string
	TraceId string
}

func passInformationAcrossSystemWithContextUsingStruct() {
	fmt.Println("> passInformationAcrossSystemWithContextUsingStruct")
	const key = "request-info"
	requestContext := RequestContext{
		UserId:  "uuid1",
		TraceId: "uuid2",
	}
	ctx := context.WithValue(context.Background(), key, requestContext)
	if id, ok := ctx.Value(key).(RequestContext); ok {
		fmt.Printf("the value is %v\n", id)
	} else {
		fmt.Println("no value in context")
	}

	// the value is {uuid1 uuid2}
}
