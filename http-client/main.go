package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const VIA_CEP = "https://viacep.com.br/ws/01001000/json"

func main() {
	fmt.Println("> http client")
	clientWithTimeout()
	doPost()
	customRequest()
	requestWithContext()
}

func clientWithTimeout() {
	fmt.Println("> clientWithTimeout")
	client := http.Client{Timeout: time.Second}
	resp, err := client.Get(VIA_CEP)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func doPost() {
	fmt.Println("> doPost")
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "John Doe"}`))
	resp, err := client.Post("https://jsonplaceholder.typicode.com/posts", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func customRequest() {
	fmt.Println("> customRequest")
	client := http.Client{}
	req, err := http.NewRequest("GET", VIA_CEP, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1l")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func requestWithContext() {
	fmt.Println("> requestWithContext")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", VIA_CEP, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
