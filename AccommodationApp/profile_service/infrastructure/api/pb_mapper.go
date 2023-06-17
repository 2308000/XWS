package api

import (
	auth "accommodation_booking/common/domain"
	pb "accommodation_booking/common/proto/profile_service"
	"accommodation_booking/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPbToAddress(pbProfile *pb.Profile) *domain.Address {
	address := &domain.Address{
		Country: pbProfile.Address.Country,
		City:    pbProfile.Address.City,
		Street:  pbProfile.Address.Street,
	}

	return address
}

func mapProfileToPb(profile *domain.Profile) *pb.Profile {
	var id *string
	primitive := profile.Id.Hex()
	id = &primitive
	pbProfile := &pb.Profile{
		Id:                    *id,
		Username:              profile.Username,
		FirstName:             profile.FirstName,
		LastName:              profile.LastName,
		Email:                 profile.Email,
		Address:               &pb.Address{Country: profile.Address.Country, City: profile.Address.City, Street: profile.Address.Street},
		DateOfBirth:           timestamppb.New(profile.DateOfBirth),
		PhoneNumber:           profile.PhoneNumber,
		Gender:                profile.Gender,
		Token:                 profile.Token,
		ReservationsCancelled: int32(profile.ReservationsCancelled),
		IsOutstanding:         profile.IsOutstanding,
	}

	return pbProfile
}

func mapPbToProfile(pbProfile *pb.Profile) *domain.Profile {
	profile := &domain.Profile{
		Id:                    getObjectId(pbProfile.Id),
		Username:              pbProfile.Username,
		FirstName:             pbProfile.FirstName,
		LastName:              pbProfile.LastName,
		FullName:              pbProfile.FirstName + pbProfile.LastName,
		Email:                 pbProfile.Email,
		Address:               domain.Address{Country: pbProfile.Address.Country, City: pbProfile.Address.City, Street: pbProfile.Address.Street},
		DateOfBirth:           pbProfile.DateOfBirth.AsTime(),
		PhoneNumber:           pbProfile.PhoneNumber,
		Gender:                pbProfile.Gender,
		Token:                 pbProfile.Token,
		ReservationsCancelled: int(pbProfile.ReservationsCancelled),
		IsOutstanding:         pbProfile.IsOutstanding,
	}

	return profile
}

func mapAuthProfileToProfile(authProfile *auth.Profile) *domain.Profile {
	profile := &domain.Profile{
		Id:                    authProfile.Id,
		Username:              authProfile.Username,
		FirstName:             authProfile.FirstName,
		LastName:              authProfile.LastName,
		FullName:              authProfile.FirstName + authProfile.LastName,
		Email:                 authProfile.Email,
		Address:               domain.Address{Country: authProfile.Address.Country, City: authProfile.Address.City, Street: authProfile.Address.Street},
		DateOfBirth:           authProfile.DateOfBirth,
		PhoneNumber:           authProfile.PhoneNumber,
		Gender:                authProfile.Gender,
		Token:                 authProfile.Token,
		ReservationsCancelled: authProfile.ReservationsCancelled,
		IsOutstanding:         authProfile.IsOutstanding,
	}

	return profile
}

func MapProfileToAuthProfile(profile *domain.Profile) *auth.Profile {
	authProfile := &auth.Profile{
		Id:          profile.Id,
		Username:    profile.Username,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		FullName:    profile.FirstName + profile.LastName,
		Email:       profile.Email,
		Address:     auth.Address(profile.Address),
		DateOfBirth: profile.DateOfBirth,
		PhoneNumber: profile.PhoneNumber,
		Gender:      profile.Gender,
		Token:       profile.Token,
	}

	return authProfile
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
