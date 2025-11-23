package main

import (
	"log"
	"product_service/internal/config"
	"product_service/internal/transport/http"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	cfg := config.LoadConfig()

	app := http.NewRouter()

	if err := app.Run(cfg.PORT); err != nil {
		panic(err)
	}
}