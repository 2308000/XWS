package api

import (
	pb "accommodation_booking/common/proto/profile_service"
	"accommodation_booking/profile_service/application"
	"accommodation_booking/profile_service/domain"
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service  *application.ProfileService
	validate *validator.Validate
}

func NewProfileHandler(service *application.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		service:  service,
		validate: domain.NewProfileValidator(),
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	profileId := request.Id
	Profile, err := handler.service.Get(ctx, profileId)
	if err != nil {
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)
	response := &pb.GetResponse{
		Profile: ProfilePb,
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
