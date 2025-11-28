package main

import (
	"product_service/internal/config"
	"product_service/internal/config/eureka_client"
	"product_service/internal/transport/http"
)

func main() {

	eurekaclient.EurekaClient()

	cfg := config.LoadConfig()

	app := http.NewRouter()

	if err := app.Run(cfg.PORT); err != nil {
		panic(err)
	}
}
