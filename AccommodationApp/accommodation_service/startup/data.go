package startup

import (
	"accommodation_booking/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:                getObjectId("62706d1b724b3da748f63fe3"),
		Host:              domain.Host{HostId: getObjectId("62706d1b624b4da648f53fe3"), Username: "host", PhoneNumber: "0667762600"},
		Name:              "soba krnjaca",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Vase Pelagica 12"},
		HasWifi:           true,
		HasFreeParking:    true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
