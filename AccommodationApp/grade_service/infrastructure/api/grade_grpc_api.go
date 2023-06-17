package api

import (
	pb "accommodation_booking/common/proto/grade_service"
	"accommodation_booking/grade_service/application"
	"accommodation_booking/grade_service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GradeHandler struct {
	pb.UnimplementedGradeServiceServer
	service *application.GradeService
}

func NewGradeHandler(service *application.GradeService) *GradeHandler {
	return &GradeHandler{
		service: service,
	}
}

func (handler *GradeHandler) Get(ctx context.Context, request *pb.GetGradeRequest) (*pb.GetGradeResponse, error) {
	gradeId := request.Id
	Grade, err := handler.service.Get(ctx, gradeId)
	if err != nil {
		return nil, err
	}
	GradePb := mapGradeToPb(Grade)
	response := &pb.GetGradeResponse{
		Grade: GradePb,
	}
	return response, nil
}

func (handler *GradeHandler) GetHostsGradedByGuest(ctx context.Context, request *pb.GetGradeRequest) (*pb.GetAllGradesResponse, error) {
	Grades, err := handler.service.GetHostsGradedByGuest(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, Grade := range Grades {
		current := mapGradeToPb(Grade)
		response.Grades = append(response.Grades, current)
	}
	return response, nil
}

func (handler *GradeHandler) GetAccommodationsGradedByGuest(ctx context.Context, request *pb.GetGradeRequest) (*pb.GetAllGradesResponse, error) {
	Grades, err := handler.service.GetAccommodationsGradedByGuest(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, Grade := range Grades {
		current := mapGradeToPb(Grade)
		response.Grades = append(response.Grades, current)
	}
	return response, nil
}

func (handler *GradeHandler) GetByGraded(ctx context.Context, request *pb.GetGradeRequest) (*pb.GetAllGradesResponse, error) {
	Grades, err := handler.service.GetByGraded(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, Grade := range Grades {
		current := mapGradeToPb(Grade)
		response.Grades = append(response.Grades, current)
	}
	return response, nil
}

func (handler *GradeHandler) GetByHost(ctx context.Context, request *pb.GetGradeRequest) (*pb.GetAllGradesResponse, error) {
	Grades, err := handler.service.GetByGraded(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, Grade := range Grades {
		current := mapGradeToPb(Grade)
		response.Grades = append(response.Grades, current)
	}
	return response, nil
}

func (handler *GradeHandler) GetAll(ctx context.Context, request *pb.GetAllGradesRequest) (*pb.GetAllGradesResponse, error) {
	grades, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, grade := range grades {
		current := mapGradeToPb(grade)
		response.Grades = append(response.Grades, current)
	}
	return response, nil
}

func (handler GradeHandler) Create(ctx context.Context, request *pb.CreateGradeRequest) (*pb.CreateGradeResponse, error) {
	guestId, err := primitive.ObjectIDFromHex(request.Grade.GuestId)
	gradedId, err := primitive.ObjectIDFromHex(request.Grade.GradedId)
	domainGrade := &domain.Grade{
		Id:       primitive.NewObjectID(),
		GuestId:  guestId,
		GradedId: gradedId,
		Grade:    request.Grade.Value,
		Date:     time.Now(),
	}
	grade, err := handler.service.Create(ctx, domainGrade)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGradeResponse{
		Grade: mapGradeToPb(grade),
	}, nil
}

func (handler GradeHandler) Update(ctx context.Context, request *pb.UpdateGradeRequest) (*pb.UpdateGradeResponse, error) {
	grade, err := handler.service.Get(ctx, request.Id)
	grade.Date = time.Now()
	grade.Grade = request.Value
	grade, err = handler.service.Update(ctx, request.Id, grade)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateGradeResponse{
		Grade: mapGradeToPb(grade),
	}, nil
}

func (handler *GradeHandler) Delete(ctx context.Context, request *pb.DeleteGradeRequest) (*pb.DeleteGradeResponse, error) {
	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGradeResponse{}, nil
}
