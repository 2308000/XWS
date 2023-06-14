package api

import (
	"accommodation_booking/accommodation_service/application"
	pb "accommodation_booking/common/proto/accommodation_service"
	"context"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
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

func (handler AccommodationHandler) Create(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	accommodation := mapPbToAccommodation(request.Accommodation)

	accommodation, err := handler.service.Create(ctx, accommodation)
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
