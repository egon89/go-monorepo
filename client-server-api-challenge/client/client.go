package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Exchange struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var exc Exchange
	err = json.Unmarshal(body, &exc)
	if err != nil {
		log.Fatal(err)
	}

	err = createFile(&exc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exchange: %v\n", exc.Bid)
}

func createFile(exchange *Exchange) error {
	f, err := os.Create("/tmp/cotacao.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	content := fmt.Sprintf("Dólar: %v", exchange.Bid)
	size, err := f.Write([]byte(content))
	if err != nil {
		return err
	}
	log.Printf("file created! Size: %d\n", size)

	return nil
}
