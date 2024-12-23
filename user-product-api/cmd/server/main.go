package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/configs"
	"github.com/egon89/go-monorepo/user-product-api/infra/database"
	"github.com/egon89/go-monorepo/user-product-api/infra/webserver/handlers"
	"github.com/egon89/go-monorepo/user-product-api/internal/entity"
	"github.com/go-chi/chi/v5"
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

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Post("/users", userHandler.Create)

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
	})

	serverPort := configs.ServerPort
	log.Printf("listening on %v", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), r)
}
