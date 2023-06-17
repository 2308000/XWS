package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port                        string
	ProfileDBHost               string
	ProfileDBPort               string
	ReservationHost             string
	ReservationPort             string
	GradeHost                   string
	GradePort                   string
	UserHost                    string
	UserPort                    string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
	CreateProfileCommandSubject string
	CreateProfileReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:                        os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost:               os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort:               os.Getenv("PROFILE_DB_PORT"),
		ReservationHost:             os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:             os.Getenv("RESERVATION_SERVICE_PORT"),
		GradeHost:                   os.Getenv("GRADE_SERVICE_HOST"),
		GradePort:                   os.Getenv("GRADE_SERVICE_PORT"),
		UserHost:                    os.Getenv("USER_SERVICE_HOST"),
		UserPort:                    os.Getenv("USER_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
		CreateProfileCommandSubject: os.Getenv("CREATE_PROFILE_COMMAND_SUBJECT"),
		CreateProfileReplySubject:   os.Getenv("CREATE_PROFILE_REPLY_SUBJECT"),
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
