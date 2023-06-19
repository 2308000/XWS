package startup

import (
	"accommodation_booking/grade_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var grades = []*domain.Grade{
	//Ocene hosta 1
	{
		Id:          getObjectId("dddddddd9999999999000000"),
		GuestId:     getObjectId("aaaaaaaa9876543210000000"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000001"),
		GuestId:     getObjectId("aaaaaaaa9876543210000001"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000002"),
		GuestId:     getObjectId("aaaaaaaa9876543210000002"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000003"),
		GuestId:     getObjectId("aaaaaaaa9876543210000003"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000004"),
		GuestId:     getObjectId("aaaaaaaa9876543210000004"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000005"),
		GuestId:     getObjectId("aaaaaaaa9876543210000005"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000006"),
		GuestId:     getObjectId("aaaaaaaa9876543210000006"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000007"),
		GuestId:     getObjectId("aaaaaaaa9876543210000007"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000008"),
		GuestId:     getObjectId("aaaaaaaa9876543210000008"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000009"),
		GuestId:     getObjectId("aaaaaaaa9876543210000009"),
		GradedId:    getObjectId("aaaaaaaa0123456789000000"),
		GradedName:  "Nikola Vukic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	//Ocene hosta 2
	{
		Id:          getObjectId("dddddddd9999999999000010"),
		GuestId:     getObjectId("aaaaaaaa9876543210000005"),
		GradedId:    getObjectId("aaaaaaaa0123456789000001"),
		GradedName:  "Jovan Gavrilovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000011"),
		GuestId:     getObjectId("aaaaaaaa9876543210000006"),
		GradedId:    getObjectId("aaaaaaaa0123456789000001"),
		GradedName:  "Jovan Gavrilovic",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000012"),
		GuestId:     getObjectId("aaaaaaaa9876543210000007"),
		GradedId:    getObjectId("aaaaaaaa0123456789000001"),
		GradedName:  "Jovan Gavrilovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000013"),
		GuestId:     getObjectId("aaaaaaaa9876543210000008"),
		GradedId:    getObjectId("aaaaaaaa0123456789000001"),
		GradedName:  "Jovan Gavrilovic",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000014"),
		GuestId:     getObjectId("aaaaaaaa9876543210000009"),
		GradedId:    getObjectId("aaaaaaaa0123456789000001"),
		GradedName:  "Jovan Gavrilovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	//Ocene hosta 3
	{
		Id:          getObjectId("dddddddd9999999999000015"),
		GuestId:     getObjectId("aaaaaaaa9876543210000005"),
		GradedId:    getObjectId("aaaaaaaa0123456789000002"),
		GradedName:  "Marija Stefanovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000016"),
		GuestId:     getObjectId("aaaaaaaa9876543210000006"),
		GradedId:    getObjectId("aaaaaaaa0123456789000002"),
		GradedName:  "Marija Stefanovic",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000017"),
		GuestId:     getObjectId("aaaaaaaa9876543210000007"),
		GradedId:    getObjectId("aaaaaaaa0123456789000002"),
		GradedName:  "Marija Stefanovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000018"),
		GuestId:     getObjectId("aaaaaaaa9876543210000008"),
		GradedId:    getObjectId("aaaaaaaa0123456789000002"),
		GradedName:  "Marija Stefanovic",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("dddddddd9999999999000019"),
		GuestId:     getObjectId("aaaaaaaa9876543210000009"),
		GradedId:    getObjectId("aaaaaaaa0123456789000002"),
		GradedName:  "Marija Stefanovic",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
