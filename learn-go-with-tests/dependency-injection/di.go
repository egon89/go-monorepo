package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Stdout implements io.Writer
	Greet(os.Stdout, "John Doe")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
