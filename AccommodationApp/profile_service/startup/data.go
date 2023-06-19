package startup

import (
	"accommodation_booking/profile_service/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var profiles = []*domain.Profile{
	{
		Id:          getObjectId("aaaaaaaa0123456789000000"),
		Username:    "host1",
		FirstName:   "Nikola",
		LastName:    "Vukic",
		FullName:    "Nikola Vukic",
		DateOfBirth: time.Date(2000, 8, 23, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "060000000",
		Email:       "host1@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Bosna i Hercegovina",
			City:    "Doboj",
			Street:  "Marsala Tita 33",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa0123456789000001"),
		Username:    "host2",
		FirstName:   "Jovan",
		LastName:    "Gavrilovic",
		FullName:    "Jovan Gavrilovic",
		DateOfBirth: time.Date(1979, 3, 17, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "060111111",
		Email:       "host2@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Novi Sad",
			Street:  "Filipa Visnjica 5",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa0123456789000002"),
		Username:    "host3",
		FirstName:   "Marija",
		LastName:    "Stefanovic",
		FullName:    "Marija Stefanovic",
		DateOfBirth: time.Date(1989, 5, 5, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "060222222",
		Email:       "host3@gmail.com",
		Gender:      "female",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Jagodina",
			Street:  "Beogradska 2",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa0123456789000003"),
		Username:    "host4",
		FirstName:   "Aleksa",
		LastName:    "Seferovic",
		FullName:    "Aleksa Seferovic",
		DateOfBirth: time.Date(1994, 3, 26, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "060333333",
		Email:       "host4@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Crna Gora",
			City:    "Ulcinj",
			Street:  "Tivatska 32",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa0123456789000004"),
		Username:    "host5",
		FirstName:   "Smiljana",
		LastName:    "Katic",
		FullName:    "Smiljana Katic",
		DateOfBirth: time.Date(2003, 8, 6, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "060444444",
		Email:       "host5@gmail.com",
		Gender:      "female",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Ecka",
			Street:  "Zrenjaninski put 87",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa9876543210000000"),
		Username:    "guest1",
		FirstName:   "Dane",
		LastName:    "Milisic",
		FullName:    "Dane Milisic",
		DateOfBirth: time.Date(2000, 4, 9, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "061000000",
		Email:       "guest1@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Hrvatska",
			City:    "Split",
			Street:  "Rivska 32",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa9876543210000001"),
		Username:    "guest2",
		FirstName:   "Stefan",
		LastName:    "Kljajic",
		FullName:    "Stefan Kljajic",
		DateOfBirth: time.Date(1995, 11, 30, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "061111111",
		Email:       "guest2@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Cacak",
			Street:  "Ibarska 89",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa9876543210000002"),
		Username:    "guest3",
		FirstName:   "Milica",
		LastName:    "Perovic",
		FullName:    "Milica Perovic",
		DateOfBirth: time.Date(2004, 3, 11, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "061222222",
		Email:       "guest3@gmail.com",
		Gender:      "female",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Krusevac",
			Street:  "Vojvode Supljikca 1",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa9876543210000003"),
		Username:    "guest4",
		FirstName:   "Aleksandra",
		LastName:    "Markovic",
		FullName:    "Aleksandra Markovic",
		DateOfBirth: time.Date(1983, 1, 20, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "061333333",
		Email:       "guest4@gmail.com",
		Gender:      "female",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Bor",
			Street:  "Djerdapska 129",
		},
		IsOutstanding: false,
	},
	{
		Id:          getObjectId("aaaaaaaa9876543210000004"),
		Username:    "guest5",
		FirstName:   "Marinko",
		LastName:    "Svetozarevic",
		FullName:    "Marinko Svetozarevic",
		DateOfBirth: time.Date(1978, 5, 6, 0, 0, 0, 0, time.UTC),
		PhoneNumber: "061444444",
		Email:       "guest4@gmail.com",
		Gender:      "male",
		Address: domain.Address{
			Country: "Srbija",
			City:    "Klek",
			Street:  "Zrenjaninska 71",
		},
		IsOutstanding: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
