package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReservationStatusType int8

const (
	Pending = iota
	Approved
	Canceled
)

type Reservation struct {
	Id                primitive.ObjectID    `bson:"_id"`
	AccommodationId   primitive.ObjectID    `bson:"accommodationId"`
	UserId            primitive.ObjectID    `bson:"userId"`
	Beginning         time.Time             `bson:"beginning"`
	Ending            time.Time             `bson:"ending"`
	Guests            int32                 `bson:"guests"`
	ReservationStatus ReservationStatusType `bson:"reservationStatus"`
}
