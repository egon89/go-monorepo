package main

import (
	"fmt"
	"golang-fiber-keycloak/api/middlewares"
	"golang-fiber-keycloak/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	initViper()

	app := fiber.New(fiber.Config{
		AppName:      "golang-fiber-keycloak",
		ServerHeader: "Fiber",
	})

	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, routes.InitProtectedRoutes)

	listenIp := viper.GetString("LISTEN_IP")
	listenPort := viper.GetString("LISTEN_PORT")

	log.Printf("will listen on %v:%v", listenIp, listenPort)

	err := app.Listen(fmt.Sprintf("%v:%v", listenIp, listenPort))
	log.Fatal(err)
}

func initViper() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}

	log.Println("viper config initialized")
}
