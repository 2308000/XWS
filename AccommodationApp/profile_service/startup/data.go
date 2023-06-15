package startup

import (
	"accommodation_booking/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var profiles = []*domain.Profile{
	{
		Id:          getObjectId("62706d1b624b3da748f63fe3"),
		Username:    "admin",
		FirstName:   "Nikola",
		LastName:    "Vukic",
		FullName:    "NikolaVukic",
		DateOfBirth: time.Time{},
		PhoneNumber: "0665762600",
		Email:       "admin@gmail.com",
		Gender:      "male",
	},
	{
		Id:          getObjectId("62706d1b623b3da748f63fa1"),
		Username:    "host",
		FirstName:   "Dane",
		LastName:    "Milisic",
		FullName:    "DaneMilisic",
		DateOfBirth: time.Time{},
		PhoneNumber: "063123123",
		Email:       "host@gmail.com",
		Gender:      "male",
	},
	{
		Id:          getObjectId("55306d1b623b3da748f63fa1"),
		Username:    "guest",
		FirstName:   "Dejan",
		LastName:    "Barcal",
		FullName:    "DejanBarcal",
		DateOfBirth: time.Time{},
		PhoneNumber: "063321321",
		Email:       "guest@gmail.com",
		Gender:      "male",
	},
	{
		Id:          getObjectId("62706d1b623b4da748f63fa1"),
		Username:    "guest",
		FirstName:   "Sergej",
		LastName:    "Madic",
		FullName:    "SergejMadic",
		DateOfBirth: time.Time{},
		PhoneNumber: "063111111",
		Email:       "guest2@gmail.com",
		Gender:      "male",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
