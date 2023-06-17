package startup

import (
	"accommodation_booking/grade_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var grades = []*domain.Grade{
	{
		Id:       getObjectId("11706d1b624b3da748f63123"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("62706d1b623b3da748f63fa1"),
		Grade:    4.9,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("22706d1b623b3da748f63223"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("62706d1b623b4da748f63bc3"),
		Grade:    4.5,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("33306d1b623b3da748f63323"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("55506d1b724b3da748f63fe3"),
		Grade:    4.8,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("44706d1b623b4da748f63423"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("77706d1b724b3da748f63fe3"),
		Grade:    4.7,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("552706d1b623b3da748f63523"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("88806d1b724b3da748f63fe3"),
		Grade:    4.3,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("66306d1b623b3da748f63623"),
		GuestId:  getObjectId("62706d1b624b3da748f63fe3"),
		GradedId: getObjectId("99906d1b724b3da748f63fe3"),
		Grade:    3.8,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("11706d1b624b3da748f63823"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("62706d1b623b3da748f63fa1"),
		Grade:    4.8,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("22706d1b623b3da748f63923"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("62706d1b623b4da748f63bc3"),
		Grade:    4.3,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("33306d1b623b3da748f63103"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("55506d1b724b3da748f63fe3"),
		Grade:    5.0,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("44706d1b623b4da748f63113"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("77706d1b724b3da748f63fe3"),
		Grade:    4.9,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("552706d1b623b3da748f63123"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("88806d1b724b3da748f63fe3"),
		Grade:    3.9,
		Date:     time.Now().AddDate(0, 0, -10),
	},
	{
		Id:       getObjectId("66306d1b623b3da748f63f133"),
		GuestId:  getObjectId("55306d1b623b3da748f63fa1"),
		GradedId: getObjectId("99906d1b724b3da748f63fe3"),
		Grade:    3.3,
		Date:     time.Now().AddDate(0, 0, -10),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
