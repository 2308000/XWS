package api

import (
	"accommodation_booking/common/auth"
	"accommodation_booking/common/domain"
	pb "accommodation_booking/common/proto/user_service"
	"accommodation_booking/user_service/application"
	"context"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service    *application.UserService
	jwtManager *auth.JWTManager
	validate   *validator.Validate
}

func NewUserHandler(service *application.UserService,
	jwtManager *auth.JWTManager) *UserHandler {
	return &UserHandler{
		service:    service,
		jwtManager: jwtManager,
		validate:   domain.NewUserValidator(),
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username := request.Username
	User, err := handler.service.Get(ctx, username)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	UserPb.Password = ""
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetById(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetResponse, error) {
	id := request.Id
	User, err := handler.service.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	UserPb.Password = ""
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUserToPb(User)
		current.Password = ""
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	request.User.Role = "user"
	mappedUser := mapPbToUser(request.User)
	if err := handler.validate.Struct(mappedUser); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	mappedUser.Password = HashPassword(mappedUser.Password)
	registeredUser, err := handler.service.Register(ctx, mappedUser, request.FirstName, request.LastName, request.Email)
	if err != nil {
		return nil, err
	}
	registeredUser.Password = ""
	return &pb.RegisterResponse{
		User: &pb.User{
			Id:       registeredUser.Id.Hex(),
			Username: registeredUser.Username,
			Role:     registeredUser.Role,
		}}, nil
}

func (handler *UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	username, err := handler.service.Update(ctx, id, request.Username)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Username: username}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := handler.service.Get(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler *UserHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := handler.service.UpdatePassword(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePasswordResponse{}, nil
}
