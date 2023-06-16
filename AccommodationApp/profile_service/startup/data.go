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
		Id:            getObjectId("62706d1b623b3da748f63fa1"),
		Username:      "host",
		FirstName:     "Dane",
		LastName:      "Milisic",
		FullName:      "DaneMilisic",
		DateOfBirth:   time.Time{},
		PhoneNumber:   "063123123",
		Email:         "host@gmail.com",
		Gender:        "male",
		IsOutstanding: true,
		Grades: []domain.Grade{
			{
				AccommodationId: getObjectId("55506d1b724b3da748f63fe3"),
				Grade:           4.8,
				Date:            time.Date(2023, time.June, 10, 0, 0, 0, 0, time.Local),
			},
		},
	},
	{
		Id:                    getObjectId("55306d1b623b3da748f63fa1"),
		Username:              "guest",
		FirstName:             "Dejan",
		LastName:              "Barcal",
		FullName:              "DejanBarcal",
		DateOfBirth:           time.Time{},
		PhoneNumber:           "063321321",
		Email:                 "guest@gmail.com",
		Gender:                "male",
		ReservationsCancelled: 2,
		Grades: []domain.Grade{
			{
				AccommodationId: getObjectId("55506d1b724b3da748f63fe3"),
				Grade:           4.8,
				Date:            time.Date(2023, time.June, 10, 0, 0, 0, 0, time.Local),
			},
			{
				AccommodationId: getObjectId("99906d1b724b3da748f63fe3"),
				Grade:           4.5,
				Date:            time.Date(2023, time.May, 10, 0, 0, 0, 0, time.Local),
			},
		},
	},
	{
		Id:            getObjectId("62706d1b623b4da748f63bc3"),
		Username:      "host",
		FirstName:     "Sergej",
		LastName:      "Madic",
		FullName:      "SergejMadic",
		DateOfBirth:   time.Time{},
		PhoneNumber:   "063111111",
		Email:         "host2@gmail.com",
		Gender:        "male",
		IsOutstanding: false,
		Grades: []domain.Grade{
			{
				AccommodationId: getObjectId("99906d1b724b3da748f63fe3"),
				Grade:           4.5,
				Date:            time.Date(2023, time.May, 10, 0, 0, 0, 0, time.Local),
			},
		},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
