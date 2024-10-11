package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog-one", blog{title: "Blog One"})
	mux.Handle("/blog-two", blog{title: "Blog Two"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

/*
	now, blog struct implements the Handler interface
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
