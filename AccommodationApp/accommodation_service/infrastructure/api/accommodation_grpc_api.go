package api

import (
	"accommodation_booking/accommodation_service/application"
	"accommodation_booking/accommodation_service/domain"
	accommodation "accommodation_booking/common/proto/accommodation_service"
	pb "accommodation_booking/common/proto/accommodation_service"
	grade "accommodation_booking/common/proto/grade_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service             *application.AccommodationService
	profileClient       profile.ProfileServiceClient
	reservationClient   reservation.ReservationServiceClient
	gradeClient         grade.GradeServiceClient
	userClient          user.UserServiceClient
	accommodationClient accommodation.AccommodationServiceClient
}

func NewAccommodationHandler(service *application.AccommodationService, profileClient profile.ProfileServiceClient, reservationClient reservation.ReservationServiceClient, gradeClient grade.GradeServiceClient, userClient user.UserServiceClient, accommodationClient accommodation.AccommodationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		service:             service,
		profileClient:       profileClient,
		reservationClient:   reservationClient,
		gradeClient:         gradeClient,
		userClient:          userClient,
		accommodationClient: accommodationClient,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetAccommodationRequest) (*pb.GetAccommodationResponse, error) {
	accommodationId := request.Id
	accommodation, err := handler.service.Get(ctx, accommodationId)
	if err != nil {
		return nil, err
	}
	accommodationPb := mapAccommodationToPb(accommodation)
	response := &pb.GetAccommodationResponse{
		Accommodation: accommodationPb,
	}
	accommodationGrades, err := handler.gradeClient.GetByGraded(ctx, &grade.GetGradeRequest{Id: accommodationId})
	totalSum := 0.0
	totalCount := 0
	for _, accommodationGrade := range accommodationGrades.Grades {
		totalSum = totalSum + float64(accommodationGrade.Grade)
		totalCount = totalCount + 1
		guestProfile, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: accommodationGrade.GuestId})
		if err != nil {
			log.Println(accommodationGrade.GuestId)
			log.Println(guestProfile)
			return nil, err
		}
		pbGrade := pb.AccommodationGrade2{
			GuestName: guestProfile.User.Username,
			Grade:     accommodationGrade.Grade,
			Date:      accommodationGrade.Date,
		}
		response.Accommodation.Grades = append(response.Accommodation.Grades, &pbGrade)
	}
	response.Accommodation.AverageAccommodationGrade = float32(totalSum / float64(totalCount))
	return response, nil
}

func (handler *AccommodationHandler) GetByHost(ctx context.Context, request *pb.GetAccommodationRequest) (*pb.GetAllAccommodationsResponse, error) {
	accommodations, err := handler.service.GetByHost(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllAccommodationsResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodationToPb(accommodation)
		accommodationGrades, err := handler.gradeClient.GetByGraded(ctx, &grade.GetGradeRequest{Id: accommodation.Id.Hex()})
		if err != nil {
			return nil, err
		}
		totalSum := 0.0
		totalCount := 0
		for _, accommodationGrade := range accommodationGrades.Grades {
			totalSum = totalSum + float64(accommodationGrade.Grade)
			totalCount = totalCount + 1
			guestProfile, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: accommodationGrade.GuestId})
			if err != nil {
				return nil, err
			}
			pbGrade := pb.AccommodationGrade2{
				GuestName: guestProfile.User.Username,
				Grade:     accommodationGrade.Grade,
				Date:      accommodationGrade.Date,
			}
			current.Grades = append(current.Grades, &pbGrade)
		}
		current.AverageAccommodationGrade = float32(totalSum / float64(totalCount))
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
	if numberOfNights == 0 {
		numberOfNights = 1
	}
	accommodations, err := handler.service.GetAllSearched(ctx, *location, int(request.NumberOfGuests))
	if err != nil {
		return nil, err
	}

	response := &pb.AccommodationSearchResponse{
		Accommodations: []*pb.AccommodationSearch{},
	}
	log.Println("Trazim slobodne smjestaje od ", request.Beginning.AsTime(), " do ", request.Ending.AsTime())
	for _, accommodation := range accommodations {
		reservations, err := handler.reservationClient.GetBetweenDates(ctx, &reservation.GetBetweenDatesRequest{Informations: &reservation.Informations{
			AccommodationId: accommodation.Id.Hex(),
			Beginning:       request.Beginning,
			Ending:          request.Ending,
		}})
		if err != nil {
			return nil, err
		}
		if len(reservations.Reservations) > 0 {
			log.Println("Smjestaj ", accommodation.Name, " ima rezervacije.")
			continue
		}
		log.Println("Smjestaj ", accommodation.Name, " nema rezervacije.")
		accommodationAvailableDateInfo, err := handler.service.GetAccommodationAvailableDatesForTimePeriod(ctx, accommodation.Id.Hex(), request.Beginning.AsTime(), request.Ending.AsTime())
		if err != nil {
			return nil, err
		}

		if len(accommodationAvailableDateInfo) == 0 {
			log.Println("Smjestaj ", accommodation.Name, " se ne moze rezervisati.")
			continue
		}
		availableDate := domain.AvailableDate{
			Beginning:       accommodationAvailableDateInfo[0].Beginning,
			Ending:          accommodationAvailableDateInfo[0].Ending,
			Price:           accommodationAvailableDateInfo[0].Price,
			IsPricePerGuest: accommodationAvailableDateInfo[0].IsPricePerGuest,
		}
		accommodation.Availability = []domain.AvailableDate{availableDate}
		accommodationPb := mapAccommodationToPb(accommodation)
		log.Println("dodao smjestaj")
		accommodationGrades, err := handler.gradeClient.GetByGraded(ctx, &grade.GetGradeRequest{Id: accommodationPb.Id})
		if err != nil {
			return nil, err
		}
		totalSum := 0.0
		totalCount := 0
		for _, accommodationGrade := range accommodationGrades.Grades {
			totalSum = totalSum + float64(accommodationGrade.Grade)
			totalCount = totalCount + 1
			guestProfile, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: accommodationGrade.GuestId})
			if err != nil {
				return nil, err
			}
			pbGrade := pb.AccommodationGrade2{
				GuestName: guestProfile.User.Username,
				Grade:     accommodationGrade.Grade,
				Date:      accommodationGrade.Date,
			}
			accommodationPb.Grades = append(accommodationPb.Grades, &pbGrade)
		}
		accommodationPb.AverageAccommodationGrade = float32(totalSum / float64(totalCount))

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

func (handler *AccommodationHandler) GetAllFiltered(ctx context.Context, request *pb.GetAllFilterRequest) (*pb.AccommodationSearchResponse, error) {
	searchResult, err := handler.GetAllSearched(ctx, request.SearchQuery)
	if err != nil {
		return nil, err
	}
	var filterSearchResult []pb.AccommodationSearch
	for _, accommodationDTO := range searchResult.Accommodations {
		if accommodationDTO.TotalPrice <= request.PriceRangeUpperBound && accommodationDTO.TotalPrice >= request.PriceRangeLowerBound {
			filterSearchResult = append(filterSearchResult, *accommodationDTO)
		}
	}
	benefits := &domain.Benefits{
		HasWifi:            request.Benefits.HasWifi,
		HasAirConditioning: request.Benefits.HasAirConditioning,
		HasWashingMachine:  request.Benefits.HasWashingMachine,
		HasBalcony:         request.Benefits.HasBalcony,
		HasKitchen:         request.Benefits.HasKitchen,
		HasBathtub:         request.Benefits.HasBathtub,
		HasFreeParking:     request.Benefits.HasFreeParking,
	}
	accommodations, err := handler.service.GetAllFiltered(ctx, *benefits)
	var indicesToKeep []int
	for _, accommodation := range accommodations {
		for i, currentDTO := range filterSearchResult {
			if currentDTO.Accommodation.Id == accommodation.Id.Hex() {
				indicesToKeep = append(indicesToKeep, i)
				continue
			}
		}
	}

	finalFilter := make([]pb.AccommodationSearch, len(indicesToKeep))
	for i, index := range indicesToKeep {
		finalFilter[i] = filterSearchResult[index]
	}

	response := pb.AccommodationSearchResponse{
		Accommodations: []*pb.AccommodationSearch{},
	}

	for i, _ := range finalFilter {
		accommodation2 := finalFilter[i].Accommodation
		if request.IsOutstandingHost == true {
			isOutstandingRes, err := handler.profileClient.IsOutstandingHost(ctx, &profile.GetRequest{Id: accommodation2.Host.HostId})
			if err != nil {
				return nil, err
			}
			if isOutstandingRes.IsOutstanding == false {
				continue
			}
		}
		finalFilter[i].Accommodation.Grades = []*pb.AccommodationGrade2{}
		accommodationGrades, err := handler.gradeClient.GetByGraded(ctx, &grade.GetGradeRequest{Id: accommodation2.Id})
		if err != nil {
			return nil, err
		}
		totalSum := 0.0
		totalCount := 0
		for _, accommodationGrade := range accommodationGrades.Grades {
			totalSum = totalSum + float64(accommodationGrade.Grade)
			totalCount = totalCount + 1
			guestProfile, err := handler.userClient.GetById(ctx, &user.GetByIdRequest{Id: accommodationGrade.GuestId})
			if err != nil {
				return nil, err
			}
			pbGrade := pb.AccommodationGrade2{
				GuestName: guestProfile.User.Username,
				Grade:     accommodationGrade.Grade,
				Date:      accommodationGrade.Date,
			}
			finalFilter[i].Accommodation.Grades = append(finalFilter[i].Accommodation.Grades, &pbGrade)
		}
		finalFilter[i].Accommodation.AverageAccommodationGrade = float32(totalSum / float64(totalCount))
		response.Accommodations = append(response.Accommodations, &finalFilter[i])
	}

	return &response, nil
}

func (handler AccommodationHandler) Create(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	if ctx.Value("userId") != request.Accommodation.HostId {
		return nil, errors.New("must be the owner to create the accommodation")
	}
	host, err := handler.profileClient.Get(ctx, &profile.GetRequest{Id: request.Accommodation.HostId})
	hostId, err := primitive.ObjectIDFromHex(host.Profile.Id)
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
		HasWifi:                       request.Accommodation.HasWifi,
		HasAirConditioning:            request.Accommodation.HasAirConditioning,
		HasFreeParking:                request.Accommodation.HasFreeParking,
		HasKitchen:                    request.Accommodation.HasKitchen,
		HasWashingMachine:             request.Accommodation.HasWashingMachine,
		HasBathtub:                    request.Accommodation.HasBathtub,
		HasBalcony:                    request.Accommodation.HasBalcony,
		Photos:                        request.Accommodation.Photos,
		MinNumberOfGuests:             int(request.Accommodation.MinNumberOfGuests),
		MaxNumberOfGuests:             int(request.Accommodation.MaxNumberOfGuests),
		IsReservationAcceptenceManual: request.Accommodation.IsReservationAcceptenceManual,
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
	if request.AvailableDate.Ending.AsTime().Before(request.AvailableDate.Ending.AsTime()) || request.AvailableDate.Ending.AsTime().Equal(request.AvailableDate.Beginning.AsTime()) {
		return nil, errors.New("cannot update availability with invalid dates")
	}
	if request.AvailableDate.Beginning.AsTime().Before(time.Now()) {
		return nil, errors.New("cannot update availability into the past")
	}
	accommodationOwnershipCheck, err := handler.accommodationClient.Get(ctx, &accommodation.GetAccommodationRequest{Id: request.AccommodationId})
	if err != nil {
		return nil, err
	}
	if accommodationOwnershipCheck.Accommodation.Host.HostId != ctx.Value("userId") {
		return nil, errors.New("must be the owner of the accommodation to update its availability")
	}

	reservationCheck, err := handler.reservationClient.GetBetweenDates(ctx, &reservation.GetBetweenDatesRequest{Informations: &reservation.Informations{
		AccommodationId: request.AccommodationId,
		Beginning:       request.AvailableDate.Beginning,
		Ending:          request.AvailableDate.Ending,
	}})
	if err != nil {
		return nil, err
	}
	if len(reservationCheck.Reservations) != 0 {
		return nil, errors.New("the given accommodation already has reservations in given time period")
	}

	accommodationId := request.AccommodationId
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
