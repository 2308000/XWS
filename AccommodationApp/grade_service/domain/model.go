package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Grade struct {
	Id          primitive.ObjectID `bson:"_id"`
	GuestId     primitive.ObjectID `bson:"guestId"`
	GradedId    primitive.ObjectID `bson:"gradedId"`
	GradedName  string             `bson:"gradedName"`
	Grade       float32            `bson:"grade"`
	Date        time.Time          `bson:"date"`
	IsHostGrade bool               `bson:"isHostGrade"`
}
