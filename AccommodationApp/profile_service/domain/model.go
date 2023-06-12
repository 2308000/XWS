package domain

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
)

type Profile struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username" validate:"username"`
	FirstName   string             `bson:"firstName" validate:"name"`
	LastName    string             `bson:"lastName" validate:"name"`
	FullName    string             `bson:"fullName"`
	Email       string             `bson:"email"`
	Address     Address            `bson:"address"`
	DateOfBirth time.Time          `bson:"dateOfBirth"`
	PhoneNumber string             `bson:"phoneNumber"`
	Gender      string             `bson:"gender"`
	Token       string             `bson:"token"`
}

type Address struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Street  string `bson:"street"`
}

func usernameValidator(fl validator.FieldLevel) bool {
	matchString, err := regexp.MatchString(`^[_a-zA-Z0-9]([._-]([._-]?)|[a-zA-Z0-9]){3,18}[_a-zA-Z0-9]$`, fl.Field().String())
	if err != nil {
		return false
	}
	return matchString
}

func nameValidator(fl validator.FieldLevel) bool {
	matchString, err := regexp.MatchString(`^[A-Z][a-z]+$`, fl.Field().String())
	if err != nil {
		return false
	}
	return matchString
}

func NewProfileValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidator)
	validate.RegisterValidation("name", nameValidator)
	return validate
}
