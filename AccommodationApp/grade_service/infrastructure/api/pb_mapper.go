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
		Id:          *id,
		GuestId:     grade.GuestId.Hex(),
		GradedName:  grade.GradedName,
		Grade:       grade.Grade,
		Date:        timestamppb.New(grade.Date),
		IsHostGrade: grade.IsHostGrade,
	}

	return pbGrade
}

func mapPbToGrade(pbGrade *pb.Grade) *domain.Grade {
	grade := &domain.Grade{
		Id:          getObjectId(pbGrade.Id),
		GuestId:     getObjectId(pbGrade.GuestId),
		GradedName:  pbGrade.GradedName,
		Grade:       pbGrade.Grade,
		Date:        pbGrade.Date.AsTime(),
		IsHostGrade: pbGrade.IsHostGrade,
	}

	return grade
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
