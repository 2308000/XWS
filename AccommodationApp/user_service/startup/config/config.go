package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("ENVF")
		}
	}
	return nil
}
