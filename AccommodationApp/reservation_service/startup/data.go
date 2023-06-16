package startup

import (
	"accommodation_booking/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var reservations = []*domain.Reservation{
	{
		Id:                getObjectId("62706d1b623ab1d748f63fa1"),
		AccommodationId:   getObjectId("55506d1b724b3da748f63fe3"),
		UserId:            getObjectId("55306d1b623b3da748f63fa1"),
		Beginning:         time.Date(2023, 7, 15, 12, 0, 0, 0, time.UTC),
		Ending:            time.Date(2023, 7, 18, 12, 0, 0, 0, time.UTC),
		Guests:            3,
		ReservationStatus: 0,
	},
	{
		Id:                getObjectId("62706d1b623ab2d748f63fa1"),
		AccommodationId:   getObjectId("55506d1b724b3da748f63fe3"),
		UserId:            getObjectId("55306d1b623b3da748f63fa1"),
		Beginning:         time.Date(2023, 7, 20, 12, 0, 0, 0, time.UTC),
		Ending:            time.Date(2023, 7, 23, 12, 0, 0, 0, time.UTC),
		Guests:            3,
		ReservationStatus: 1,
	},
	{
		Id:                getObjectId("62706d1b623ab3d748f63fa1"),
		AccommodationId:   getObjectId("55506d1b724b3da748f63fe3"),
		UserId:            getObjectId("55306d1b623b3da748f63fa1"),
		Beginning:         time.Date(2023, 7, 25, 12, 0, 0, 0, time.UTC),
		Ending:            time.Date(2023, 7, 28, 12, 0, 0, 0, time.UTC),
		Guests:            3,
		ReservationStatus: 2,
	},
	{
		Id:                getObjectId("62706d1b623ab3d781f63fa1"),
		AccommodationId:   getObjectId("55506d1b724b3da748f63fe3"),
		UserId:            getObjectId("55306d1b623b3da748f63fa1"),
		Beginning:         time.Now(),
		Ending:            time.Now().Add(3),
		Guests:            3,
		ReservationStatus: 0,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
