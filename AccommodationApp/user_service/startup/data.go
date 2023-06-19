package startup

import (
	auth "accommodation_booking/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*auth.User{
	{
		Id:       getObjectId("aaaaaaaa0123456789000000"),
		Username: "host1",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "guest",
	},
	{
		Id:       getObjectId("aaaaaaaa0123456789000001"),
		Username: "host2",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa0123456789000002"),
		Username: "host3",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "guest",
	},
	{
		Id:       getObjectId("aaaaaaaa0123456789000003"),
		Username: "host4",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "guest",
	},
	{
		Id:       getObjectId("aaaaaaaa0123456789000004"),
		Username: "host5",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa9876543210000000"),
		Username: "guest1",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa9876543210000001"),
		Username: "guest2",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa9876543210000002"),
		Username: "guest3",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa9876543210000003"),
		Username: "guest4",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
	{
		Id:       getObjectId("aaaaaaaa9876543210000004"),
		Username: "guest5",
		Password: "$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i",
		Role:     "host",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
