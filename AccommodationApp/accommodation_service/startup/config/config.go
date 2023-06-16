package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port                        string
	AccommodationDBHost         string
	AccommodationDBPort         string
	ProfileHost                 string
	ProfilePort                 string
	ReservationHost             string
	ReservationPort             string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                        os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:         os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:         os.Getenv("ACCOMMODATION_DB_PORT"),
		ProfileHost:                 os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:                 os.Getenv("PROFILE_SERVICE_PORT"),
		ReservationHost:             os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:             os.Getenv("RESERVATION_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
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
