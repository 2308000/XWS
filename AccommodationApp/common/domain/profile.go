package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username" validate:"username"`
	FirstName string             `bson:"firstName" validate:"name"`
	LastName  string             `bson:"lastName" validate:"name"`
	FullName  string             `bson:"fullName"`
	Address   Address            `bson:"address"`
}

type Address struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Street  string `bson:"street"`
}
