package api

import (
	accommodation "accommodation_booking/common/proto/accommodation_service"
	profile "accommodation_booking/common/proto/profile_service"
	pb "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	"accommodation_booking/reservation_service/application"
	"accommodation_booking/reservation_service/domain"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service             *application.ReservationService
	userClient          user.UserServiceClient
	accommodationClient accommodation.AccommodationServiceClient
	profileClient       profile.ProfileServiceClient
}

func NewReservationHandler(service *application.ReservationService, userClient user.UserServiceClient,
	accommodationClient accommodation.AccommodationServiceClient, profileClient profile.ProfileServiceClient) *ReservationHandler {
	return &ReservationHandler{
		service:             service,
		userClient:          userClient,
		accommodationClient: accommodationClient,
		profileClient:       profileClient,
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
	foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
	if err != nil {
		return nil, err
	}
	accommodationDetails := &pb.AccommodationDetails{
		Id:   foundAccommodation.Accommodation.Id,
		Name: foundAccommodation.Accommodation.Name,
	}
	ReservationPb := &pb.ReservationOut{
		Id:                Reservation.Id.Hex(),
		Accommodation:     accommodationDetails,
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
	Reservations, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllReservationsResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:       foundUser.User.Id,
			Username: foundUser.User.Username,
		}
		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) GetBetweenDates(ctx context.Context, request *pb.GetBetweenDatesRequest) (*pb.GetBetweenDatesResponse, error) {
	log.Println("usao u grpc")
	Reservations, err := handler.service.GetBetweenDates(ctx, request.Informations.Beginning.AsTime(), request.Informations.Ending.AsTime(), request.Informations.AccommodationId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetBetweenDatesResponse{
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

func (handler *ReservationHandler) GetUsersReservations(ctx context.Context, request *pb.GetUsersReservationsRequest) (*pb.GetUsersReservationsResponse, error) {
	Reservations, err := handler.service.GetForUser(ctx, request.UserId, "")
	if err != nil {
		return nil, err
	}
	response := &pb.GetUsersReservationsResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:       foundUser.User.Id,
			Username: foundUser.User.Username,
		}
		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetByHost(ctx context.Context, request *pb.GetByHostRequest) (*pb.GetByHostResponse, error) {
	allReservations, err := handler.service.GetPending(ctx)
	if err != nil {
		return nil, err
	}
	Reservations := []*domain.Reservation{}
	for _, res := range allReservations {
		acc, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: res.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		if acc.Accommodation.Host.HostId == ctx.Value("userId").(string) {
			Reservations = append(Reservations, res)
		}
	}
	response := &pb.GetByHostResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		foundProfile, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:                  foundUser.User.Id,
			Username:            foundUser.User.Username,
			CancellationCounter: foundProfile.Profile.ReservationsCancelled,
		}

		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetByHostPending(ctx context.Context, request *pb.GetByHostInternRequest) (*pb.GetByHostInternResponse, error) {
	allReservations, err := handler.service.GetPending(ctx)
	if err != nil {
		return nil, err
	}
	Reservations := []*domain.Reservation{}
	for _, res := range allReservations {
		acc, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: res.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		if acc.Accommodation.Host.HostId == request.Id {
			Reservations = append(Reservations, res)
		}
	}
	response := &pb.GetByHostInternResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		foundProfile, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:                  foundUser.User.Id,
			Username:            foundUser.User.Username,
			CancellationCounter: foundProfile.Profile.ReservationsCancelled,
		}

		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetByHostCanceled(ctx context.Context, request *pb.GetByHostInternRequest) (*pb.GetByHostInternResponse, error) {
	allReservations, err := handler.service.GetCanceled(ctx)
	if err != nil {
		return nil, err
	}
	Reservations := []*domain.Reservation{}
	for _, res := range allReservations {
		acc, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: res.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		if acc.Accommodation.Host.HostId == request.Id {
			Reservations = append(Reservations, res)
		}
	}
	response := &pb.GetByHostInternResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		foundProfile, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:                  foundUser.User.Id,
			Username:            foundUser.User.Username,
			CancellationCounter: foundProfile.Profile.ReservationsCancelled,
		}

		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetByHostApproved(ctx context.Context, request *pb.GetByHostInternRequest) (*pb.GetByHostInternResponse, error) {
	allReservations, err := handler.service.GetApproved(ctx)
	if err != nil {
		return nil, err
	}
	Reservations := []*domain.Reservation{}
	for _, res := range allReservations {
		acc, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: res.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		if acc.Accommodation.Host.HostId == request.Id {
			Reservations = append(Reservations, res)
		}
	}
	response := &pb.GetByHostInternResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		foundProfile, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:                  foundUser.User.Id,
			Username:            foundUser.User.Username,
			CancellationCounter: foundProfile.Profile.ReservationsCancelled,
		}

		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetByHostRejected(ctx context.Context, request *pb.GetByHostInternRequest) (*pb.GetByHostInternResponse, error) {
	allReservations, err := handler.service.GetRejected(ctx)
	if err != nil {
		return nil, err
	}
	Reservations := []*domain.Reservation{}
	for _, res := range allReservations {
		acc, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: res.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		if acc.Accommodation.Host.HostId == request.Id {
			Reservations = append(Reservations, res)
		}
	}
	response := &pb.GetByHostInternResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		foundProfile, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:                  foundUser.User.Id,
			Username:            foundUser.User.Username,
			CancellationCounter: foundProfile.Profile.ReservationsCancelled,
		}

		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
			Beginning:         timestamppb.New(Reservation.Beginning),
			Ending:            timestamppb.New(Reservation.Ending),
			Guests:            Reservation.Guests,
			ReservationStatus: int32(Reservation.ReservationStatus),
		}
		response.Reservations = append(response.Reservations, current)

	}
	return response, nil
}

func (handler *ReservationHandler) GetMyReservations(ctx context.Context, request *pb.GetMyReservationsRequest) (*pb.GetMyReservationsResponse, error) {
	userId := ctx.Value("userId").(string)
	Reservations, err := handler.service.GetForUser(ctx, userId, request.ResType)
	if err != nil {
		return nil, err
	}
	response := &pb.GetMyReservationsResponse{
		Reservations: []*pb.ReservationOut{},
	}
	for _, Reservation := range Reservations {
		foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: Reservation.UserId.Hex()})
		if err != nil {
			return nil, err
		}
		userDetails := &pb.UserDetails{
			Id:       foundUser.User.Id,
			Username: foundUser.User.Username,
		}
		foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: Reservation.AccommodationId.Hex()})
		if err != nil {
			return nil, err
		}
		accommodationDetails := &pb.AccommodationDetails{
			Id:   foundAccommodation.Accommodation.Id,
			Name: foundAccommodation.Accommodation.Name,
		}
		current := &pb.ReservationOut{
			Id:                Reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
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
	checkAvailability, err := handler.accommodationClient.GetAccommodationAvailableDatesForTimePeriod(ctx, &accommodation.AccommodationTimePeriodRequest{
		AccommodationId: request.Reservation.AccommodationId,
		Beginning:       request.Reservation.Beginning,
		Ending:          request.Reservation.Ending,
	})
	if len(checkAvailability.AvailableDates) == 0 {
		return nil, errors.New("host set accommodation as not available during selected period")
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
	selectedAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: accommodationId.Hex()})
	if err != nil {
		return nil, err
	}
	//log.Println(request.Reservation.Guests, selectedAccommodation.Accommodation.MinNumberOfGuests, request.Reservation.Guests, selectedAccommodation.Accommodation.MaxNumberOfGuests)
	if request.Reservation.Guests < selectedAccommodation.Accommodation.MinNumberOfGuests || request.Reservation.Guests > selectedAccommodation.Accommodation.MaxNumberOfGuests {
		return nil, errors.New("guest number restriction violation")
	}
	status := 0
	if selectedAccommodation.Accommodation.IsReservationAcceptenceManual {
		status = 1
	}
	reservation := &domain.Reservation{
		Id:                primitive.NewObjectID(),
		AccommodationId:   accommodationId,
		UserId:            userId,
		Beginning:         request.Reservation.Beginning.AsTime(),
		Ending:            request.Reservation.Ending.AsTime(),
		Guests:            request.Reservation.Guests,
		ReservationStatus: domain.ReservationStatusType(status),
	}
	err = handler.service.Create(ctx, reservation)
	if err != nil {
		return nil, err
	}
	foundUser, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: userId.Hex()})
	if err != nil {
		return nil, err
	}
	userDetails := &pb.UserDetails{
		Id:       foundUser.User.Id,
		Username: foundUser.User.Username,
	}
	foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: accommodationId.Hex()})
	if err != nil {
		return nil, err
	}
	accommodationDetails := &pb.AccommodationDetails{
		Id:   foundAccommodation.Accommodation.Id,
		Name: foundAccommodation.Accommodation.Name,
	}
	return &pb.CreateReservationResponse{
		Reservation: &pb.ReservationOut{
			Id:                reservation.Id.Hex(),
			Accommodation:     accommodationDetails,
			User:              userDetails,
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
		ReservationStatus: domain.ReservationStatusType(1),
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
	reservation, err := handler.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: reservation.AccommodationId.Hex()})
	if err != nil {
		return nil, err
	}
	if foundAccommodation.Accommodation.Host.HostId != ctx.Value("userId").(string) {
		return nil, errors.New("you cannot approve this reservation")
	}
	err = handler.service.Approve(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ApproveReservationResponse{}, nil
}

func (handler *ReservationHandler) Cancel(ctx context.Context, request *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	reservation, err := handler.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if reservation.UserId.Hex() != ctx.Value("userId").(string) {
		return nil, errors.New("you are not allowed cancel this reservation")
	}
	tomorrow := time.Now().Add(1)
	if reservation.Beginning.Before(tomorrow) {
		return nil, errors.New("you cannot cancel a reservation if there is less than a day left until it starts")
	}
	if reservation.ReservationStatus == 0 {
		err = handler.service.Delete(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.CancelReservationResponse{}, nil
	} else {
		err = handler.service.Cancel(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		response, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: ctx.Value("userId").(string)})
		if err != nil {
			return nil, err
		}
		_, err = handler.profileClient.Update(ctx, &profile.UpdateRequest{
			Id: response.Profile.Id,
			Profile: &profile.Profile{
				Id:                    response.Profile.Id,
				Username:              response.Profile.Username,
				FirstName:             response.Profile.FirstName,
				LastName:              response.Profile.LastName,
				Email:                 response.Profile.Email,
				Address:               response.Profile.Address,
				DateOfBirth:           response.Profile.DateOfBirth,
				PhoneNumber:           response.Profile.PhoneNumber,
				Gender:                response.Profile.Gender,
				Token:                 response.Profile.Token,
				ReservationsCancelled: response.Profile.ReservationsCancelled + 1,
				IsOutstanding:         response.Profile.IsOutstanding,
			},
		})
		if err != nil {
			return nil, err
		}
		return &pb.CancelReservationResponse{}, nil
	}
}

func (handler *ReservationHandler) Reject(ctx context.Context, request *pb.RejectReservationRequest) (*pb.RejectReservationResponse, error) {
	reservation, err := handler.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	foundAccommodation, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: reservation.AccommodationId.Hex()})
	if err != nil {
		return nil, err
	}
	log.Println(foundAccommodation.Accommodation.Host.HostId)
	log.Println(ctx.Value("userId").(string))
	if foundAccommodation.Accommodation.Host.HostId != ctx.Value("userId").(string) {
		return nil, errors.New("you cannot reject this reservation")
	}
	err = handler.service.Reject(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RejectReservationResponse{}, nil
}
