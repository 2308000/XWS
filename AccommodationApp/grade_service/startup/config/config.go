package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port              string
	GradeDBHost       string
	GradeDBPort       string
	ProfileHost       string
	ProfilePort       string
	ReservationHost   string
	ReservationPort   string
	AccommodationHost string
	AccommodationPort string
	NatsHost          string
	NatsPort          string
	NatsUser          string
	NatsPass          string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:              os.Getenv("GRADE_SERVICE_PORT"),
		GradeDBHost:       os.Getenv("GRADE_DB_HOST"),
		GradeDBPort:       os.Getenv("GRADE_DB_PORT"),
		ProfileHost:       os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:       os.Getenv("PROFILE_SERVICE_PORT"),
		ReservationHost:   os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		NatsHost:          os.Getenv("NATS_HOST"),
		NatsPort:          os.Getenv("NATS_PORT"),
		NatsUser:          os.Getenv("NATS_USER"),
		NatsPass:          os.Getenv("NATS_PASS"),
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
