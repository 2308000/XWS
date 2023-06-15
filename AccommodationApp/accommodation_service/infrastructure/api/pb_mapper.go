package api

import (
	"accommodation_booking/accommodation_service/domain"
	pb "accommodation_booking/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return &pb.Accommodation{
		Id:                 accommodation.Id.Hex(),
		Host:               mapHostToPb(accommodation),
		Name:               accommodation.Name,
		Location:           mapLocationToPb(accommodation),
		HasWifi:            accommodation.HasWifi,
		HasAirConditioning: accommodation.HasAirConditioning,
		HasFreeParking:     accommodation.HasFreeParking,
		HasKitchen:         accommodation.HasKitchen,
		HasWashingMachine:  accommodation.HasWashingMachine,
		HasBathtub:         accommodation.HasBathtub,
		HasBalcony:         accommodation.HasBalcony,
		Photos:             accommodation.Photos,
		MinNumberOfGuests:  int32(accommodation.MinNumberOfGuests),
		MaxNumberOfGuests:  int32(accommodation.MaxNumberOfGuests),
	}
}

// Function to return a domain.Accommodation from a pb.Accommodation
func mapPbToAccommodation(pbAccommodation *pb.Accommodation) *domain.Accommodation {
	return &domain.Accommodation{
		Id:                 getObjectId(pbAccommodation.Id),
		Host:               domain.Host{HostId: getObjectId(pbAccommodation.Host.HostId), Username: pbAccommodation.Host.Username, PhoneNumber: pbAccommodation.Host.PhoneNumber},
		Name:               pbAccommodation.Name,
		Location:           domain.Location{Country: pbAccommodation.Location.Country, City: pbAccommodation.Location.City, Street: pbAccommodation.Location.Street},
		HasWifi:            pbAccommodation.HasWifi,
		HasAirConditioning: pbAccommodation.HasAirConditioning,
		HasFreeParking:     pbAccommodation.HasFreeParking,
		HasKitchen:         pbAccommodation.HasKitchen,
		HasWashingMachine:  pbAccommodation.HasWashingMachine,
		HasBathtub:         pbAccommodation.HasBathtub,
		HasBalcony:         pbAccommodation.HasBalcony,
		Photos:             pbAccommodation.Photos,
		MinNumberOfGuests:  int(pbAccommodation.MinNumberOfGuests),
		MaxNumberOfGuests:  int(pbAccommodation.MaxNumberOfGuests),
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
