package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// os.Stdout implements io.Writer
	Greet(os.Stdout, "John Doe\n")

	fmt.Println("starting server on 5001...")

	log.Fatal(http.ListenAndServe(":5001",
		http.HandlerFunc(MyGreeterHandler)))
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter implements io.Writer
	Greet(w, "Jane Doe")
}
