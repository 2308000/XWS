package api

import (
	"accommodation_booking/accommodation_service/application"
	"accommodation_booking/accommodation_service/domain"
	pb "accommodation_booking/common/proto/accommodation_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service           *application.AccommodationService
	profileClient     profile.ProfileServiceClient
	reservationClient reservation.ReservationServiceClient
}

func NewAccommodationHandler(service *application.AccommodationService, profileClient profile.ProfileServiceClient, reservationClient reservation.ReservationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		service:           service,
		profileClient:     profileClient,
		reservationClient: reservationClient,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetAccommodationRequest) (*pb.GetAccommodationResponse, error) {
	accommodationId := request.Id
	Accommodation, err := handler.service.Get(ctx, accommodationId)
	if err != nil {
		return nil, err
	}
	AccommodationPb := mapAccommodationToPb(Accommodation)
	response := &pb.GetAccommodationResponse{
		Accommodation: AccommodationPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetByHost(ctx context.Context, request *pb.GetAccommodationRequest) (*pb.GetAllAccommodationsResponse, error) {
	Accommodations, err := handler.service.GetByHost(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllAccommodationsResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, Accommodation := range Accommodations {
		current := mapAccommodationToPb(Accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllAccommodationsRequest) (*pb.GetAllAccommodationsResponse, error) {
	accommodations, err := handler.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllAccommodationsResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodationToPb(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllSearched(ctx context.Context, request *pb.AccommodationSearchRequest) (*pb.AccommodationSearchResponse, error) {
	location := &domain.Location{
		Country: request.Location.Country,
		City:    request.Location.City,
		Street:  request.Location.Street,
	}
	numberOfNights := int(request.Ending.AsTime().Sub(request.Beginning.AsTime()).Hours() / 24)
	accommodations, indices, err := handler.service.GetAllSearched(ctx, *location, request.Beginning.AsTime(), request.Ending.AsTime(), int(request.NumberOfGuests))

	if err != nil {
		return nil, err
	}

	response := &pb.AccommodationSearchResponse{
		Accommodations: []*pb.AccommodationSearch{},
	}
	for i, accommodation := range accommodations {
		accommodation.Availability = []domain.AvailableDate{accommodation.Availability[indices[i]]}
		accommodationPb := mapAccommodationToPb(accommodation)
		pricePerNight := float64(accommodation.Availability[0].Price)
		totalPrice := float64(numberOfNights) * pricePerNight
		if accommodation.Availability[0].IsPricePerGuest {
			totalPrice = totalPrice * float64(request.NumberOfGuests)
		}
		response.Accommodations = append(response.Accommodations, &pb.AccommodationSearch{
			Accommodation: accommodationPb,
			TotalPrice:    float32(totalPrice),
			PricePerNight: float32(pricePerNight),
		})
	}
	return response, err
}

func (handler *AccommodationHandler) GetAllFiltered(ctx context.Context, request *pb.GetAllFilterRequest) (*pb.GetAllAccommodationsResponse, error) {
	benefits := &domain.Benefits{
		HasWifi:            request.Benefits.HasWifi,
		HasAirConditioning: request.Benefits.HasAirConditioning,
		HasWashingMachine:  request.Benefits.HasWashingMachine,
		HasBalcony:         request.Benefits.HasBalcony,
		HasKitchen:         request.Benefits.HasKitchen,
		HasBathtub:         request.Benefits.HasBathtub,
		HasFreeParking:     request.Benefits.HasFreeParking,
	}
	Accommodations, err := handler.service.GetAllFiltered(ctx, request.PriceRangeLowerBound, request.PriceRangeUpperBound, *benefits, request.IsOutstandingHost)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllAccommodationsResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, Accommodation := range Accommodations {
		current := mapAccommodationToPb(Accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler AccommodationHandler) Create(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	host, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: request.Accommodation.HostId})
	hostId, err := primitive.ObjectIDFromHex(host.Profile.Id)
	/*log.Println(ctx.Value("userId").(string))
	log.Println(ctx.Value("username").(string))
	if ctx.Value("userId").(string) != hostId.Hex() {
		return nil, errors.New("must be the host to add new accommodation")
	}*/
	hostInfo := &domain.Host{
		HostId:        hostId,
		Username:      host.Profile.Username,
		PhoneNumber:   host.Profile.PhoneNumber,
		IsOutstanding: true,
	}
	accommodationInfo := &domain.Accommodation{
		Id:   primitive.NewObjectID(),
		Host: *hostInfo,
		Name: request.Accommodation.Name,
		Location: domain.Location{
			Country: request.Accommodation.Location.Country,
			City:    request.Accommodation.Location.City,
			Street:  request.Accommodation.Location.Street,
		},
		HasWifi:            request.Accommodation.HasWifi,
		HasAirConditioning: request.Accommodation.HasAirConditioning,
		HasFreeParking:     request.Accommodation.HasFreeParking,
		HasKitchen:         request.Accommodation.HasKitchen,
		HasWashingMachine:  request.Accommodation.HasWashingMachine,
		HasBathtub:         request.Accommodation.HasBathtub,
		HasBalcony:         request.Accommodation.HasBalcony,
		Photos:             request.Accommodation.Photos,
		MinNumberOfGuests:  int(request.Accommodation.MinNumberOfGuests),
		MaxNumberOfGuests:  int(request.Accommodation.MaxNumberOfGuests),
	}
	accommodation, err := handler.service.Create(ctx, accommodationInfo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccommodationResponse{
		Accommodation: mapAccommodationToPb(accommodation),
	}, nil
}

func (handler AccommodationHandler) Update(ctx context.Context, request *pb.UpdateAccommodationRequest) (*pb.UpdateAccommodationResponse, error) {
	accommodation := mapPbToAccommodation(request.Accommodation)
	accommodation, err := handler.service.Update(ctx, request.Id, accommodation)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccommodationResponse{
		Accommodation: mapAccommodationToPb(accommodation),
	}, nil
}

func (handler *AccommodationHandler) UpdateAvailability(ctx context.Context, request *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	accommodationId := request.AccommodationId
	//reservations, err := handler.reservationClient.
	availableDate := domain.AvailableDate{
		Beginning:       request.AvailableDate.Beginning.AsTime(),
		Ending:          request.AvailableDate.Ending.AsTime(),
		Price:           request.AvailableDate.Price,
		IsPricePerGuest: request.AvailableDate.IsPricePerGuest,
	}
	accommodation, err := handler.service.UpdateAvailability(ctx, accommodationId, availableDate)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAvailabilityResponse{Accommodation: mapAccommodationToPb(accommodation)}, err
}

func (handler *AccommodationHandler) GetAccommodationAvailableDatesForTimePeriod(ctx context.Context, request *pb.AccommodationTimePeriodRequest) (*pb.AccommodationAvailableDatesForTimePeriodResponse, error) {
	accommodation, err := handler.service.Get(ctx, request.AccommodationId)
	if err != nil {
		return nil, err
	}

	availableDates, err := handler.service.GetAccommodationAvailableDatesForTimePeriod(ctx, request.AccommodationId, request.Beginning.AsTime(), request.Ending.AsTime())
	response := &pb.AccommodationAvailableDatesForTimePeriodResponse{
		AccommodationId:   request.AccommodationId,
		AccommodationName: accommodation.Name,
	}

	for _, availableDate := range availableDates {
		availableDatePb := pb.AvailableDate{
			Beginning:       timestamppb.New(availableDate.Beginning),
			Ending:          timestamppb.New(availableDate.Ending),
			Price:           availableDate.Price,
			IsPricePerGuest: availableDate.IsPricePerGuest,
		}
		response.AvailableDates = append(response.AvailableDates, &availableDatePb)
	}

	return response, err
}

func (handler *AccommodationHandler) Delete(ctx context.Context, request *pb.DeleteAccommodationRequest) (*pb.DeleteAccommodationResponse, error) {
	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccommodationResponse{}, nil
}
