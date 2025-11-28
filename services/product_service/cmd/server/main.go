package main

import (
	"log"
	"product_service/internal/config"
	eurekaclient "product_service/internal/config/eureka_client"
	"product_service/internal/transport/http"
)

func main() {

	cfg := config.LoadConfig()

	// Register with Eureka
	if err := eurekaclient.RegisterEurekaClient(cfg.PORT); err != nil {
		log.Printf("Warning: Failed to register with Eureka: %v", err)
	}

	app := http.NewRouter()

	if err := app.Run(cfg.PORT); err != nil {
		panic(err)
	}
}