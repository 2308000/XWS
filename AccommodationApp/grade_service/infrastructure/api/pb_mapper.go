package api

import (
	pb "accommodation_booking/common/proto/grade_service"
	"accommodation_booking/grade_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapGradeToPb(grade *domain.Grade) *pb.Grade {
	var id *string
	primitive := grade.Id.Hex()
	id = &primitive
	pbGrade := &pb.Grade{
		Id:       *id,
		GuestId:  grade.GuestId.Hex(),
		GradedId: grade.GradedId.Hex(),
		Grade:    grade.Grade,
		Date:     timestamppb.New(grade.Date),
	}

	return pbGrade
}

func mapPbToGrade(pbGrade *pb.Grade) *domain.Grade {
	grade := &domain.Grade{
		Id:       getObjectId(pbGrade.Id),
		GuestId:  getObjectId(pbGrade.GuestId),
		GradedId: getObjectId(pbGrade.GradedId),
		Grade:    pbGrade.Grade,
		Date:     pbGrade.Date.AsTime(),
	}

	return grade
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
