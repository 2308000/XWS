package api

import (
	accommodation "accommodation_booking/common/proto/accommodation_service"
	pb "accommodation_booking/common/proto/grade_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	"accommodation_booking/grade_service/application"
	"accommodation_booking/grade_service/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GradeHandler struct {
	pb.UnimplementedGradeServiceServer
	service             *application.GradeService
	profileClient       profile.ProfileServiceClient
	accommodationClient accommodation.AccommodationServiceClient
	reservationClient   reservation.ReservationServiceClient
}

func NewGradeHandler(service *application.GradeService, profileClient profile.ProfileServiceClient, accommodationClient accommodation.AccommodationServiceClient, reservationClient reservation.ReservationServiceClient) *GradeHandler {
	return &GradeHandler{
		service:             service,
		profileClient:       profileClient,
		accommodationClient: accommodationClient,
		reservationClient:   reservationClient,
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
	grades, err := handler.service.GetByGraded(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllGradesResponse{
		Grades: []*pb.Grade{},
	}
	for _, grade := range grades {
		current := mapGradeToPb(grade)
		guestInfo, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: grade.GuestId.Hex()})
		if err != nil {
			return nil, err
		}
		current.GradedName = guestInfo.Profile.Username
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
	if ctx.Value("userId").(string) != request.Grade.GuestId {
		return nil, errors.New("you cannot grade for others")
	}

	guestId, err := primitive.ObjectIDFromHex(request.Grade.GuestId)
	gradedId, err := primitive.ObjectIDFromHex(request.Grade.GradedId)
	gradedName := ""
	if request.Grade.IsHostGrade == true {
		accommodationInfo, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: request.Grade.GradedId})
		if err != nil {
			return nil, err
		}
		gradedProfileInfo, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: accommodationInfo.Accommodation.Host.HostId})
		if err != nil {
			return nil, err
		}
		gradedId, err = primitive.ObjectIDFromHex(accommodationInfo.Accommodation.Host.HostId)
		if err != nil {
			return nil, err
		}
		gradedName = gradedProfileInfo.Profile.FirstName + " " + gradedProfileInfo.Profile.LastName
	} else {
		gradedAccommodationInfo, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: request.Grade.GradedId})
		if err != nil {
			return nil, err
		}
		gradedName = gradedAccommodationInfo.Accommodation.Name
	}
	domainGrade := &domain.Grade{
		Id:          primitive.NewObjectID(),
		GuestId:     guestId,
		GradedId:    gradedId,
		GradedName:  gradedName,
		Grade:       request.Grade.Value,
		Date:        time.Now(),
		IsHostGrade: request.Grade.IsHostGrade,
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
	gradeInfo, err := handler.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if gradeInfo.GuestId.Hex() != ctx.Value("userId").(string) {
		return nil, errors.New("you can only update your own grades")
	}

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
	gradeInfo, err := handler.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if gradeInfo.GuestId.Hex() != ctx.Value("userId").(string) {
		return nil, errors.New("you can only delete your own grades")
	}
	err = handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGradeResponse{}, nil
}
