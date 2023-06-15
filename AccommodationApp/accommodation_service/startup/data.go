package startup

import (
	"accommodation_booking/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:                getObjectId("55506d1b724b3da748f63fe3"),
		Host:              domain.Host{HostId: getObjectId("62706d1b623b3da748f63fa1"), Username: "host", PhoneNumber: "0667762600", IsOutstanding: true},
		Name:              "soba krnjaca",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Vase Pelagica 12"},
		HasWifi:           true,
		HasFreeParking:    true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
	},
	{
		Id:                 getObjectId("77706d1b724b3da748f63fe3"),
		Host:               domain.Host{HostId: getObjectId("62706d1b623b3da748f63fa1"), Username: "host", PhoneNumber: "0667762600", IsOutstanding: true},
		Name:               "stan na dan LUX LUX LUX TOP TOP",
		Location:           domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Bulevar Oslobodjenja 43"},
		HasWifi:            true,
		HasFreeParking:     true,
		HasBalcony:         true,
		HasAirConditioning: true,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  7,
	},
	{
		Id:                getObjectId("88806d1b724b3da748f63fe3"),
		Host:              domain.Host{HostId: getObjectId("62706d1b623b4da748f63bc3"), Username: "sega", PhoneNumber: "0667762611", IsOutstanding: false},
		Name:              "vila sa bazenom",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Mite Ruzica 8"},
		HasWifi:           true,
		HasKitchen:        true,
		HasBalcony:        true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
	},
	{
		Id:                getObjectId("99906d1b724b3da748f63fe3"),
		Host:              domain.Host{HostId: getObjectId("62706d1b623b4da748f63bc3"), Username: "sega", PhoneNumber: "0667762611", IsOutstanding: false},
		Name:              "penthaus neki",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Bele Njive 18"},
		HasWifi:           true,
		HasFreeParking:    true,
		HasWashingMachine: true,
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
