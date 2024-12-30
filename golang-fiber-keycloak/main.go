package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	initViper()

	listenIp := viper.GetString("LISTEN_IP")
	listenPort := viper.GetString("LISTEN_PORT")

	log.Printf("will listen on %v:%v", listenIp, listenPort)
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
