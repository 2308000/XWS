package startup

import (
	"accommodation_booking/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var accommodations = []*domain.Accommodation{
	{
		Id: getObjectId("55506d1b724b3da748f63fe3"),
		Host: domain.Host{
			HostId:   getObjectId("62706d1b623b3da748f63fa1"),
			Username: "host", PhoneNumber: "0667762600",
			IsOutstanding: true},
		Name:              "soba krnjaca",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Vase Pelagica 12"},
		HasWifi:           true,
		HasFreeParking:    true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 1, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 15, 23, 59, 59, 999, time.Local),
				Price:           40,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.June, 16, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 25, 23, 59, 59, 999, time.Local),
				Price:           50,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.June, 29, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.July, 05, 23, 59, 59, 999, time.Local),
				Price:           60,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.July, 10, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.July, 31, 23, 59, 59, 999, time.Local),
				Price:           60,
				IsPricePerGuest: true,
			},
		},
		IsReservationAcceptenceManual: false,
	},
	{
		Id: getObjectId("77706d1b724b3da748f63fe3"),
		Host: domain.Host{
			HostId:   getObjectId("62706d1b623b3da748f63fa1"),
			Username: "host", PhoneNumber: "0667762600",
			IsOutstanding: true},
		Name:               "stan na dan LUX LUX LUX TOP TOP",
		Location:           domain.Location{Country: "Srbija", City: "Beograd", Street: "Bulevar Oslobodjenja 43"},
		HasWifi:            true,
		HasFreeParking:     true,
		HasBalcony:         true,
		HasAirConditioning: true,
		MinNumberOfGuests:  2,
		MaxNumberOfGuests:  7,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 10, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 15, 23, 59, 59, 999, time.Local),
				Price:           20,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.June, 16, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 20, 23, 59, 59, 999, time.Local),
				Price:           45,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.June, 29, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.July, 9, 23, 59, 59, 999, time.Local),
				Price:           60,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 10, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.July, 31, 23, 59, 59, 999, time.Local),
				Price:           80,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: true,
	},
	{
		Id: getObjectId("88806d1b724b3da748f63fe3"),
		Host: domain.Host{
			HostId:        getObjectId("62706d1b623b4da748f63bc3"),
			Username:      "sega",
			PhoneNumber:   "0667762611",
			IsOutstanding: false},
		Name:              "vila sa bazenom",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Mite Ruzica 8"},
		HasWifi:           true,
		HasKitchen:        true,
		HasBalcony:        true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 15, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 25, 23, 59, 59, 999, time.Local),
				Price:           20,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.July, 1, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.July, 31, 23, 59, 59, 999, time.Local),
				Price:           45,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.August, 1, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.August, 9, 23, 59, 59, 999, time.Local),
				Price:           60,
				IsPricePerGuest: false,
			},
			{
				Beginning:       time.Date(2023, time.August, 15, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.August, 31, 23, 59, 59, 999, time.Local),
				Price:           55,
				IsPricePerGuest: false,
			},
		},
		IsReservationAcceptenceManual: false,
	},
	{
		Id: getObjectId("99906d1b724b3da748f63fe3"),
		Host: domain.Host{
			HostId:   getObjectId("62706d1b623b4da748f63bc3"),
			Username: "sega", PhoneNumber: "0667762611",
			IsOutstanding: false},
		Name:              "penthaus neki",
		Location:          domain.Location{Country: "Srbija", City: "Novi Sad", Street: "Bele Njive 18"},
		HasWifi:           true,
		HasFreeParking:    true,
		HasWashingMachine: true,
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 7,
		Availability: []domain.AvailableDate{
			{
				Beginning:       time.Date(2023, time.June, 20, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.June, 25, 23, 59, 59, 999, time.Local),
				Price:           20,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.August, 1, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.August, 9, 23, 59, 59, 999, time.Local),
				Price:           60,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.August, 15, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.August, 31, 23, 59, 59, 999, time.Local),
				Price:           55,
				IsPricePerGuest: true,
			},
			{
				Beginning:       time.Date(2023, time.September, 1, 0, 0, 0, 0, time.Local),
				Ending:          time.Date(2023, time.September, 31, 23, 59, 59, 999, time.Local),
				Price:           40,
				IsPricePerGuest: true,
			},
		},
		IsReservationAcceptenceManual: true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
