package startup

import (
	auth "accommodation_booking/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*auth.User{
	{
		Id:       getObjectId("62706d1b624b3da748f63fe3"),
		Username: "admin",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "admin",
	},
	{
		Id:       getObjectId("62706d1b624b4da648f53fe3"),
		Username: "host",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("62706d1b623b3da748f63fa1"),
		Username: "guest",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "guest",
	},
	{
		Id:       getObjectId("55306d1b623b3da748f63fa1"),
		Username: "dane",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "guest",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
