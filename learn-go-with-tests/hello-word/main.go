package main

import "fmt"

const hello = "Hello"

func main() {
	fmt.Println(Hello("John Doe"))
}

func Hello(name string) string {
	return fmt.Sprintf("%s %s", hello, name)
}
