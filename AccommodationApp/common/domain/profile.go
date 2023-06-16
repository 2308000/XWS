package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Profile struct {
	Id                    primitive.ObjectID `bson:"_id"`
	Username              string             `bson:"username" validate:"username"`
	FirstName             string             `bson:"firstName" validate:"name"`
	LastName              string             `bson:"lastName" validate:"name"`
	FullName              string             `bson:"fullName"`
	Email                 string             `bson:"email"`
	Address               Address            `bson:"address"`
	DateOfBirth           time.Time          `bson:"dateOfBirth"`
	PhoneNumber           string             `bson:"phoneNumber"`
	Gender                string             `bson:"gender"`
	Token                 string             `bson:"token"`
	ReservationsCancelled int                `bson:"reservationsCancelled"`
	IsOutstanding         bool               `bson:"isOutstanding"`
	Grades                []Grade            `bson:"grades"`
}

type Grade struct {
	AccommodationId primitive.ObjectID `bson:"accommodationId"`
	Grade           float64            `bson:"grade"`
	Date            time.Time          `bson:"date"`
}

type Address struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Street  string `bson:"street"`
}
