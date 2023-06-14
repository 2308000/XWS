package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Host               Host               `bson:"host"`
	Name               string             `bson:"name"`
	Location           Location           `bson:"location"`
	HasWifi            bool               `bson:"hasWifi"`
	HasAirConditioning bool               `bson:"hasAirConditioning"`
	HasFreeParking     bool               `bson:"hasFreeParking"`
	HasKitchen         bool               `bson:"hasKitchen"`
	HasWashingMachine  bool               `bson:"hasWashingMachine"`
	HasBathtub         bool               `bson:"hasBathtub"`
	HasBalcony         bool               `bson:"hasBalcony"`
	Photos             [][]byte           `bson:"photos"`
	MinNumberOfGuests  int                `bson:"minGuests"`
	MaxNumberOfGuests  int                `bson:"maxGuests"`
}

type Host struct {
	HostId      primitive.ObjectID `bson:"hostId"`
	Username    string             `bson:"username"`
	PhoneNumber string             `bson:"phoneNumber"`
}

type Location struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Street  string `bson:"street"`
}
