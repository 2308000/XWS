package startup

import (
	"accommodation_booking/grade_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var grades = []*domain.Grade{
	{
		Id:          getObjectId("11706d1b624b3da748f63123"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("62706d1b623b3da748f63fa1"),
		GradedName:  "Dane Milisic",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("22706d1b623b3da748f63223"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("62706d1b623b4da748f63bc3"),
		GradedName:  "Sergej Madic",
		Grade:       4.5,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("33306d1b623b3da748f63323"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("55506d1b724b3da748f63fe3"),
		GradedName:  "soba krnjaca",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("44706d1b623b4da748f63423"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("77706d1b724b3da748f63fe3"),
		GradedName:  "stan na dan LUX LUX LUX TOP TOP",
		Grade:       4.7,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("552706d1b623b3da748f63523"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("88806d1b724b3da748f63fe3"),
		GradedName:  "vila sa bazenom",
		Grade:       4.3,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("66306d1b623b3da748f63623"),
		GuestId:     getObjectId("62706d1b624b3da748f63fe3"),
		GradedId:    getObjectId("99906d1b724b3da748f63fe3"),
		GradedName:  "penthaus neki",
		Grade:       3.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("11706d1b624b3da748f63823"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("62706d1b623b3da748f63fa1"),
		GradedName:  "Dane Milisic",
		Grade:       4.8,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("22706d1b623b3da748f63923"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("62706d1b623b4da748f63bc3"),
		GradedName:  "Sergej Madic",
		Grade:       4.3,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: true,
	},
	{
		Id:          getObjectId("33306d1b623b3da748f63103"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("55506d1b724b3da748f63fe3"),
		GradedName:  "soba krnjaca",
		Grade:       5.0,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("44706d1b623b4da748f63113"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("77706d1b724b3da748f63fe3"),
		GradedName:  "stan na dan LUX LUX LUX TOP TOP",
		Grade:       4.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("552706d1b623b3da748f63123"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("88806d1b724b3da748f63fe3"),
		GradedName:  "vila sa bazenom",
		Grade:       3.9,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
	{
		Id:          getObjectId("66306d1b623b3da748f63f133"),
		GuestId:     getObjectId("55306d1b623b3da748f63fa1"),
		GradedId:    getObjectId("99906d1b724b3da748f63fe3"),
		GradedName:  "penthaus neki",
		Grade:       3.3,
		Date:        time.Now().AddDate(0, 0, -10),
		IsHostGrade: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
