package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port     string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:     os.Getenv("GATEWAY_PORT"),
		UserHost: os.Getenv("USER_SERVICE_HOST"),
		UserPort: os.Getenv("USER_SERVICE_PORT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}
