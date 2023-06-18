package api

import (
	accommodation "accommodation_booking/common/proto/accommodation_service"
	grade "accommodation_booking/common/proto/grade_service"
	pb "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	"accommodation_booking/profile_service/application"
	"accommodation_booking/profile_service/domain"
	"context"
	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service             *application.ProfileService
	validate            *validator.Validate
	gradeClient         grade.GradeServiceClient
	reservationClient   reservation.ReservationServiceClient
	userClient          user.UserServiceClient
	accommodationClient accommodation.AccommodationServiceClient
}

func NewProfileHandler(service *application.ProfileService, reservationClient reservation.ReservationServiceClient, gradeClient grade.GradeServiceClient, userClient user.UserServiceClient, accommodationClient accommodation.AccommodationServiceClient) *ProfileHandler {
	return &ProfileHandler{
		service:             service,
		validate:            domain.NewProfileValidator(),
		reservationClient:   reservationClient,
		gradeClient:         gradeClient,
		userClient:          userClient,
		accommodationClient: accommodationClient,
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	profileId := request.Id
	profile, err := handler.service.Get(ctx, profileId)
	userInfo, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}
	profilePb := &pb.Profile{
		Id:        request.Id,
		Username:  profile.Username,
		FirstName: profile.Email,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Address: &pb.Address{
			Country: profile.Address.Country,
			City:    profile.Address.City,
			Street:  profile.Address.Street,
		},
		DateOfBirth:           timestamppb.New(profile.DateOfBirth),
		PhoneNumber:           profile.PhoneNumber,
		Gender:                profile.Gender,
		Token:                 profile.Token,
		ReservationsCancelled: -1,
		IsOutstanding:         false,
		HostGrades:            []*pb.HostGrade{},
		AccommodationGrades:   []*pb.AccommodationGrade{},
		AverageHostGrade:      -1,
	}
	role := userInfo.User.Role
	if role == "guest" {
		profilePb.ReservationsCancelled = int32(profile.ReservationsCancelled)
		hostGrades, err := handler.gradeClient.GetHostsGradedByGuest(ctx, &grade.GetGradeRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}
		for _, hostGrade := range hostGrades.Grades {
			pbGrade := pb.HostGrade{
				HostName: hostGrade.GradedName,
				Grade:    hostGrade.Grade,
				Date:     hostGrade.Date,
			}
			profilePb.HostGrades = append(profilePb.HostGrades, &pbGrade)
		}
		accommodationGrades, err := handler.gradeClient.GetAccommodationsGradedByGuest(ctx, &grade.GetGradeRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}
		for _, accommodationGrade := range accommodationGrades.Grades {
			pbGrade := pb.AccommodationGrade{
				AccommodationName: accommodationGrade.GradedName,
				Grade:             accommodationGrade.Grade,
				Date:              accommodationGrade.Date,
			}
			profilePb.AccommodationGrades = append(profilePb.AccommodationGrades, &pbGrade)
		}
	} else {
		hostGrades, err := handler.gradeClient.GetByGraded(ctx, &grade.GetGradeRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}
		totalSum := 0.0
		numOfGrades := 0
		for _, hostGrade := range hostGrades.Grades {
			guestName, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: hostGrade.GuestId})
			if err != nil {
				return nil, err
			}
			totalSum = totalSum + float64(hostGrade.Grade)
			numOfGrades = numOfGrades + 1
			pbGrade := pb.HostGrade{
				HostName: guestName.User.Username,
				Grade:    hostGrade.Grade,
				Date:     hostGrade.Date,
			}
			profilePb.HostGrades = append(profilePb.HostGrades, &pbGrade)
		}
		averageGrade := totalSum / float64(numOfGrades)
		profilePb.AverageHostGrade = float32(averageGrade)
		if averageGrade > 4.7 {
			isOutstandingRes, err := handler.IsOutstandingHost(ctx, &pb.GetRequest{Id: request.Id})
			if err != nil {
				return nil, err
			}
			profilePb.IsOutstanding = isOutstandingRes.IsOutstanding
		}
	}
	response := &pb.GetResponse{
		Profile: profilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) IsOutstandingHost(ctx context.Context, request *pb.GetRequest) (*pb.IsOutstandingResponse, error) {
	hostsAccommodations, err := handler.accommodationClient.GetByHost(ctx, &accommodation.GetAccommodationRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}
	totalReservations := 0
	totalNumberOfReservedDays := 0.0
	for _, hostAccommodation := range hostsAccommodations.Accommodations {
		numberOfReservationsForAccommodation, err := handler.reservationClient.GetBetweenDates(ctx, &reservation.GetBetweenDatesRequest{Informations: &reservation.Informations{
			AccommodationId: hostAccommodation.Id,
			Beginning:       timestamppb.New(time.Now().AddDate(-10, 0, 0)),
			Ending:          timestamppb.New(time.Now()),
		}})
		if err != nil {
			return nil, err
		}
		if len(numberOfReservationsForAccommodation.Reservations) > 0 {
			totalReservations = totalReservations + len(numberOfReservationsForAccommodation.Reservations)
			for _, accReservation := range numberOfReservationsForAccommodation.Reservations {
				numberOfReservedDays := accReservation.Ending.AsTime().Sub(accReservation.Beginning.AsTime()).Hours() / 24
				totalNumberOfReservedDays = totalNumberOfReservedDays + numberOfReservedDays
			}
		}
	}

	response := &pb.IsOutstandingResponse{IsOutstanding: false}
	if totalReservations >= 1 && totalNumberOfReservedDays > 1 {
		cancelledReservations, err := handler.reservationClient.GetByHostCanceled(ctx, &reservation.GetByHostInternRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}
		cancellationPercentage := 0.0
		cancellationPercentage = (float64(len(cancelledReservations.Reservations)) / float64(totalReservations)) * 100
		if cancellationPercentage < 20 {
			response.IsOutstanding = true
		}
	}

	return response, nil
}
func (handler *ProfileHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Profiles, err := handler.service.GetAll(ctx, strings.ReplaceAll(request.Search, " ", ""))
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Profiles: []*pb.Profile{},
	}
	for _, Profile := range Profiles {
		current := mapProfileToPb(Profile)
		response.Profiles = append(response.Profiles, current)
	}
	return response, nil
}

func (handler ProfileHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	empty := ""
	value := &empty
	request.Profile.PhoneNumber = *value
	request.Profile.DateOfBirth = timestamppb.New(time.Now())
	request.Profile.Gender = *value

	profile := mapPbToProfile(request.Profile)
	if err := handler.validate.Struct(profile); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	err := handler.service.Create(ctx, profile)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler ProfileHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Update(ctx, request.Id, profile)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler *ProfileHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (handler *ProfileHandler) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if ctx.Value("userId").(string) != request.Id {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}
	token, err := handler.service.GenerateToken(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{
		Token: token,
	}, nil
}

func (handler *ProfileHandler) GetByToken(ctx context.Context, request *pb.GetByTokenRequest) (*pb.GetByTokenResponse, error) {
	Profile, err := handler.service.GetByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)
	response := &pb.GetByTokenResponse{
		Profile: ProfilePb,
	}
	return response, nil
}
