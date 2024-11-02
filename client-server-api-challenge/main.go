package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *sql.DB
}

type ExchangeSymbol string

const (
	USDBRL ExchangeSymbol = "USDBRL"
	USDEUR ExchangeSymbol = "USDEUR"
)

type Exchange struct {
	Code       string
	Bid        float64
	High       float64
	Low        float64
	VarBid     float64
	PctChange  float64
	Ask        float64
	Timestamp  int64
	CreateDate time.Time
}

type ExchangeCodeAwesomeResponse struct {
	Code       string `json:"code"`
	Bid        string `json:"bid"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ExchangeAwesomeResponse map[string]ExchangeCodeAwesomeResponse

func main() {
	log.Println("starting client-server-api-challenge")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	dbPath := os.Getenv("DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("database %v connected", dbPath)
	app := &App{DB: db}
	log.Println("server is listening on port 8080...")
	http.HandleFunc("/", app.GetExchangePriceHandler)
	http.ListenAndServe(":8080", nil)
}

func (app *App) GetExchangePriceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()
	exchangeCodeResponse, err := GetAwesomeExchange(USDBRL)
	if err != nil {
		log.Fatalf("exchange integration failed: %v", err)
	}
	exchange, err := exchangeCodeResponse.mapToExchange()
	if err != nil {
		log.Fatalf("map to exchange failed: %v", err)
	}
	err = app.insertExchange(ctx, &exchange)
	if err != nil {
		log.Fatalf("error to persist exchange: %v %v", exchange.Code, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchange)
}

func GetAwesomeExchange(symbol ExchangeSymbol) (*ExchangeCodeAwesomeResponse, error) {
	symbolStr := awesomePathMap(symbol)
	url := "https://economia.awesomeapi.com.br/json/last/" + symbolStr
	log.Printf("url: %v", url)
	resp, err := http.Get(url)
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
	exchangeCodeResponse, ok := exchangeResponse[string(symbol)]
	if !ok {
		return nil, fmt.Errorf("no exchange for %v", symbol)
	}

	return &exchangeCodeResponse, nil
}

func (raw *ExchangeCodeAwesomeResponse) mapToExchange() (Exchange, error) {
	var exchange Exchange
	var err error
	exchange.Code = raw.Code
	if exchange.Bid, err = strconv.ParseFloat(raw.Bid, 64); err != nil {
		return exchange, fmt.Errorf("error parsing Bid: %w", err)
	}
	if exchange.High, err = strconv.ParseFloat(raw.High, 64); err != nil {
		return exchange, fmt.Errorf("error parsing High: %w", err)
	}
	if exchange.Low, err = strconv.ParseFloat(raw.Low, 64); err != nil {
		return exchange, fmt.Errorf("error parsing Low: %w", err)
	}
	if exchange.VarBid, err = strconv.ParseFloat(raw.VarBid, 64); err != nil {
		return exchange, fmt.Errorf("error parsing VarBid: %w", err)
	}
	if exchange.PctChange, err = strconv.ParseFloat(raw.PctChange, 64); err != nil {
		return exchange, fmt.Errorf("error parsing PctChange: %w", err)
	}
	if exchange.Ask, err = strconv.ParseFloat(raw.Ask, 64); err != nil {
		return exchange, fmt.Errorf("error parsing Ask: %w", err)
	}
	if exchange.Timestamp, err = strconv.ParseInt(raw.Timestamp, 10, 64); err != nil {
		return exchange, fmt.Errorf("error parsing Timestamp: %w", err)
	}
	if exchange.CreateDate, err = time.Parse(time.DateTime, raw.CreateDate); err != nil {
		return exchange, fmt.Errorf("error parsing CreateDate: %w", err)
	}

	return exchange, nil
}

func awesomePathMap(symbol ExchangeSymbol) string {
	pathMap := map[ExchangeSymbol]string{
		USDBRL: "USD-BRL",
	}
	path, ok := pathMap[symbol]
	if !ok {
		log.Fatalf("no awesome path for %v", symbol)
	}

	return path
}

func (app *App) insertExchange(ctx context.Context, exchange *Exchange) error {
	query := `
    INSERT INTO exchange (code, bid, high, low, var_bid, pct_change, ask, timestamp, create_date)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := app.DB.ExecContext(ctx, query,
		exchange.Code, exchange.Bid, exchange.High, exchange.Low, exchange.VarBid,
		exchange.PctChange, exchange.Ask, exchange.Timestamp, exchange.CreateDate)

	return err
}
