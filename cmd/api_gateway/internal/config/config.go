package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	AuthServiceAddress string
	ProductServiceAddress string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		ServerAddress: os.Getenv("API_GATEWAY_ADDRESS"),
		AuthServiceAddress: os.Getenv("AUTH_SERVICE_ADDRESS"),
		ProductServiceAddress: os.Getenv("PRODUCT_SERVICE_ADDRESS"),
	}

	if config.ServerAddress == "" {
		config.ServerAddress = ":8080"
	}
	if config.AuthServiceAddress == "" {
		config.AuthServiceAddress = "localhost:5000"
	}

	return config, nil
}