package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/configs"
	"github.com/egon89/go-monorepo/user-product-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println("starting...")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dbPath := fmt.Sprintf("%v/%v", configs.DBHost, configs.DBName)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("database %v connected", dbPath)

	db.AutoMigrate(&entity.User{})

	serverPort := configs.ServerPort
	log.Printf("listening on %v", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
}
