package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/configs"
)

func main() {
	log.Println("starting...")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	serverPort := configs.ServerPort
	log.Printf("Listening on %v", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
}
