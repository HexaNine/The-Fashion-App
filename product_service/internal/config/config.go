package config

import (
	"log"
	"os"
	s "product_service/pkg"
)

type Config struct {
	PORT string
	DB_URL string
}

func LoadConfig() *Config{
	
    port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	if port == "" {
		port = ":8080"
	} else if !s.HasPrefixConlon(port) {
		port = ":" + port
	}

	if s.IsEmpty(dbURL) {
		log.Fatal("error : DB_URL are not provided")
	}

	return &Config{PORT: port, DB_URL: dbURL}
} 