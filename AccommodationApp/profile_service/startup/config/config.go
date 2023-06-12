package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:          os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost: os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort: os.Getenv("PROFILE_DB_PORT"),
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
