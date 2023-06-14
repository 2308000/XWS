package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port                        string
	UserDBHost                  string
	UserDBPort                  string
	ProfileHost                 string
	ProfilePort                 string
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
		Port:                        os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:                  os.Getenv("USER_DB_HOST"),
		UserDBPort:                  os.Getenv("USER_DB_PORT"),
		ProfileHost:                 os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:                 os.Getenv("PROFILE_SERVICE_PORT"),
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
			log.Fatal("ENVF")
		}
	}
	return nil
}
