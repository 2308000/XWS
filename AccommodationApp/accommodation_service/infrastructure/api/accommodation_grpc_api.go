package api

import (
	"accommodation_booking/accommodation_service/application"
	"accommodation_booking/accommodation_service/domain"
	pb "accommodation_booking/common/proto/accommodation_service"
	profile "accommodation_booking/common/proto/profile_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service       *application.AccommodationService
	profileClient profile.ProfileServiceClient
}

func NewAccommodationHandler(service *application.AccommodationService, profileClient profile.ProfileServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		service:       service,
		profileClient: profileClient,
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
	Accommodations, err := handler.service.GetAll(ctx)
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

func (handler *AccommodationHandler) Delete(ctx context.Context, request *pb.DeleteAccommodationRequest) (*pb.DeleteAccommodationResponse, error) {
	err := handler.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccommodationResponse{}, nil
}
