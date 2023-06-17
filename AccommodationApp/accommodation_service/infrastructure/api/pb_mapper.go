package api

import (
	"accommodation_booking/accommodation_service/domain"
	pb "accommodation_booking/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapHostToPb(accommodation *domain.Accommodation) *pb.HostRes {
	return &pb.HostRes{
		HostId:        accommodation.Host.HostId.Hex(),
		Username:      accommodation.Host.Username,
		PhoneNumber:   accommodation.Host.PhoneNumber,
		IsOutstanding: accommodation.Host.IsOutstanding,
	}
}

func mapLocationToPb(accommodation *domain.Accommodation) *pb.Location {
	return &pb.Location{
		Country: accommodation.Location.Country,
		City:    accommodation.Location.City,
		Street:  accommodation.Location.Street,
	}
}

func mapAccommodationToPb(accommodation *domain.Accommodation) *pb.Accommodation {
	pbAccommodation := &pb.Accommodation{
		Id:                            accommodation.Id.Hex(),
		Host:                          mapHostToPb(accommodation),
		Name:                          accommodation.Name,
		Location:                      mapLocationToPb(accommodation),
		HasWifi:                       accommodation.HasWifi,
		HasAirConditioning:            accommodation.HasAirConditioning,
		HasFreeParking:                accommodation.HasFreeParking,
		HasKitchen:                    accommodation.HasKitchen,
		HasWashingMachine:             accommodation.HasWashingMachine,
		HasBathtub:                    accommodation.HasBathtub,
		HasBalcony:                    accommodation.HasBalcony,
		Photos:                        accommodation.Photos,
		MinNumberOfGuests:             int32(accommodation.MinNumberOfGuests),
		MaxNumberOfGuests:             int32(accommodation.MaxNumberOfGuests),
		IsReservationAcceptenceManual: accommodation.IsReservationAcceptenceManual,
		Grades:                        []*pb.AccommodationGrade2{},
	}

	for _, availableDate := range accommodation.Availability {
		availableDatePb := pb.AvailableDate{
			Beginning:       timestamppb.New(availableDate.Beginning),
			Ending:          timestamppb.New(availableDate.Ending),
			Price:           availableDate.Price,
			IsPricePerGuest: availableDate.IsPricePerGuest,
		}
		pbAccommodation.Availability = append(pbAccommodation.Availability, &availableDatePb)
	}

	return pbAccommodation
}

// Function to return a domain.Accommodation from a pb.Accommodation
func mapPbToAccommodation(pbAccommodation *pb.Accommodation) *domain.Accommodation {
	accommodation := &domain.Accommodation{
		Id:                            getObjectId(pbAccommodation.Id),
		Host:                          domain.Host{HostId: getObjectId(pbAccommodation.Host.HostId), Username: pbAccommodation.Host.Username, PhoneNumber: pbAccommodation.Host.PhoneNumber},
		Name:                          pbAccommodation.Name,
		Location:                      domain.Location{Country: pbAccommodation.Location.Country, City: pbAccommodation.Location.City, Street: pbAccommodation.Location.Street},
		HasWifi:                       pbAccommodation.HasWifi,
		HasAirConditioning:            pbAccommodation.HasAirConditioning,
		HasFreeParking:                pbAccommodation.HasFreeParking,
		HasKitchen:                    pbAccommodation.HasKitchen,
		HasWashingMachine:             pbAccommodation.HasWashingMachine,
		HasBathtub:                    pbAccommodation.HasBathtub,
		HasBalcony:                    pbAccommodation.HasBalcony,
		Photos:                        pbAccommodation.Photos,
		MinNumberOfGuests:             int(pbAccommodation.MinNumberOfGuests),
		MaxNumberOfGuests:             int(pbAccommodation.MaxNumberOfGuests),
		IsReservationAcceptenceManual: pbAccommodation.IsReservationAcceptenceManual,
	}

	for _, pbAvailableDate := range pbAccommodation.Availability {
		availableDate := domain.AvailableDate{
			Beginning:       pbAvailableDate.Beginning.AsTime(),
			Ending:          pbAvailableDate.Ending.AsTime(),
			Price:           pbAvailableDate.Price,
			IsPricePerGuest: pbAvailableDate.IsPricePerGuest,
		}
		accommodation.Availability = append(accommodation.Availability, availableDate)
	}

	return accommodation
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
