package api

import (
	pb "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	"accommodation_booking/reservation_service/application"
	"accommodation_booking/reservation_service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service    *application.ReservationService
	userClient user.UserServiceClient
}

func NewReservationHandler(service *application.ReservationService, userClient user.UserServiceClient) *ReservationHandler {
	return &ReservationHandler{
		service:    service,
		userClient: userClient,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	reservationId := request.Id
	Reservation, err := handler.service.Get(ctx, reservationId)
	if err != nil {
		return nil, err
	}
	foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
	if err != nil {
		return nil, err
	}
	userDetails := &pb.UserDetails{
		Id:       foundUser.User.Id,
		Username: foundUser.User.Username,
	}
	ReservationPb := &pb.ReservationOut{
		Id:                Reservation.Id.Hex(),
		Accommodation:     nil,
		User:              userDetails,
		Beginning:         timestamppb.New(Reservation.Beginning),
		Ending:            timestamppb.New(Reservation.Ending),
		Guests:            Reservation.Guests,
		ReservationStatus: int32(Reservation.ReservationStatus),
	}
	response := &pb.GetReservationResponse{
		Reservation: ReservationPb,
	}
	return response, nil
}

func (handler *ReservationHandler) GetAll(ctx context.Context, request *pb.GetAllReservationsRequest) (*pb.GetAllReservationsResponse, error) {
	Reservations, err := handler.service.GetAll(ctx, strings.ReplaceAll(request.Search, " ", ""))
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllReservationsResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     nil,
			User:              nil,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler ReservationHandler) Create(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	accommodationId, err := primitive.ObjectIDFromHex(request.Reservation.AccommodationId)
	if err != nil {
		return nil, err
	}
	rawUserId := ctx.Value("userId").(string)
	userId, err := primitive.ObjectIDFromHex(rawUserId)
	if err != nil {
		return nil, err
	}
	reservation := &domain.Reservation{
		AccommodationId:   accommodationId,
		UserId:            userId,
		Beginning:         request.Reservation.Beginning.AsTime(),
		Ending:            request.Reservation.Ending.AsTime(),
		Guests:            request.Reservation.Guests,
		ReservationStatus: domain.ReservationStatusType(request.Reservation.ReservationStatus),
	}
	err = handler.service.Create(ctx, reservation)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReservationResponse{
		Reservation: &pb.ReservationOut{
			Id:                reservation.Id.Hex(),
			Accommodation:     nil,
			User:              nil,
			Beginning:         timestamppb.New(reservation.Beginning),
			Ending:            timestamppb.New(reservation.Ending),
			Guests:            reservation.Guests,
			ReservationStatus: int32(reservation.ReservationStatus),
		},
	}, nil
}

func (handler ReservationHandler) Update(ctx context.Context, request *pb.UpdateReservationRequest) (*pb.UpdateReservationResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Reservation.Id)
	if err != nil {
		return nil, err
	}
	accommodationId, err := primitive.ObjectIDFromHex(request.Reservation.AccommodationId)
	if err != nil {
		return nil, err
	}
	rawUserId := ctx.Value("userId").(string)
	userId, err := primitive.ObjectIDFromHex(rawUserId)
	if err != nil {
		return nil, err
	}
	reservation := &domain.Reservation{
		Id:                id,
		AccommodationId:   accommodationId,
		UserId:            userId,
		Beginning:         request.Reservation.Beginning.AsTime(),
		Ending:            request.Reservation.Ending.AsTime(),
		Guests:            request.Reservation.Guests,
		ReservationStatus: domain.ReservationStatusType(request.Reservation.ReservationStatus),
	}
	err = handler.service.Update(ctx, request.Id, reservation)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateReservationResponse{
		Reservation: &pb.ReservationOut{
			Id:                reservation.Id.Hex(),
			Accommodation:     nil,
			User:              nil,
			Beginning:         timestamppb.New(reservation.Beginning),
			Ending:            timestamppb.New(reservation.Ending),
			Guests:            reservation.Guests,
			ReservationStatus: int32(reservation.ReservationStatus),
		},
	}, nil
}

func (handler *ReservationHandler) Delete(ctx context.Context, request *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReservationResponse{}, nil
}

func (handler *ReservationHandler) Approve(ctx context.Context, request *pb.ApproveReservationRequest) (*pb.ApproveReservationResponse, error) {
	err := handler.service.Approve(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ApproveReservationResponse{}, nil
}

func (handler *ReservationHandler) Cancel(ctx context.Context, request *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	err := handler.service.Cancel(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CancelReservationResponse{}, nil
}
