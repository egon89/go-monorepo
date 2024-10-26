package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ExchangeSymbol string

const (
	USDBRL ExchangeSymbol = "USDBRL"
	USDEUR ExchangeSymbol = "USDEUR"
)

type Exchange struct {
	Code string
	Bid  float64
}

type ExchangeCodeAwesomeResponse struct {
	Code string `json:"code"`
	Bid  string `json:"bid"`
}

type ExchangeAwesomeResponse map[string]ExchangeCodeAwesomeResponse

func main() {
	fmt.Println("> client-server-api-challenge")
	fmt.Println("Server is listening on port 8080...")
	http.HandleFunc("/", GetExchangePrice)
	http.ListenAndServe(":8080", nil)
}

func GetExchangePrice(w http.ResponseWriter, r *http.Request) {
	exchangeResponse, err := GetAwesomeExchange()
	if err != nil {
		log.Fatalf("exchange integration failed %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchangeResponse)
}

func GetAwesomeExchange() (ExchangeAwesomeResponse, error) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var exchangeResponse ExchangeAwesomeResponse
	err = json.Unmarshal(body, &exchangeResponse)
	if err != nil {
		return nil, err
	}

	return exchangeResponse, nil
}
