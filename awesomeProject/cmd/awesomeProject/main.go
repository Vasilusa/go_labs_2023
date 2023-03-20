package main

import (
	"awesomeProject/internal/app/dsn"
	app2 "awesomeProject/internal/pkg/app"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Application start!")

	_ = godotenv.Load() // .env file
	toOpen := dsn.FromEnv()
	fmt.Println("toOpen = ", toOpen)
	app, err := app2.NewApplication(toOpen)

	if err != nil {
		log.Panic(err)
	}

	app.StartServer()

	log.Println("Application terminated!")
}
