package application

import (
	auth "accommodation_booking/common/domain"
	"accommodation_booking/profile_service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileService struct {
	store        domain.ProfileStore
	orchestrator *UpdateProfileOrchestrator
}

func NewProfileService(store domain.ProfileStore, orchestrator *UpdateProfileOrchestrator) *ProfileService {
	return &ProfileService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *ProfileService) Get(ctx context.Context, profileId string) (*domain.Profile, error) {
	return service.store.Get(ctx, profileId)
}

func (service *ProfileService) GetAll(ctx context.Context, search string) ([]*domain.Profile, error) {
	return service.store.GetAll(ctx, search)
}

func (service *ProfileService) Create(ctx context.Context, profile *domain.Profile) error {
	return service.store.Create(ctx, profile)
}

func (service *ProfileService) RollbackUpdate(ctx context.Context, profile *domain.Profile) error {
	return service.store.Update(ctx, profile.Id.Hex(), profile)
}

func (service *ProfileService) Update(ctx context.Context, profileId string, profile *domain.Profile) error {
	oldProfile, err := service.Get(ctx, profileId)
	if err != nil {
		return err
	}
	err = service.store.Update(ctx, profileId, profile)
	if err != nil {
		return err
	}
	newProfile := &auth.Profile{
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
	oldAuthProfile := &auth.Profile{
		Id:          oldProfile.Id,
		Username:    oldProfile.Username,
		FirstName:   oldProfile.FirstName,
		LastName:    oldProfile.LastName,
		FullName:    oldProfile.FirstName + oldProfile.LastName,
		Email:       oldProfile.Email,
		Address:     auth.Address{Country: oldProfile.Address.Country, City: oldProfile.Address.City, Street: oldProfile.Address.Street},
		DateOfBirth: oldProfile.DateOfBirth,
		PhoneNumber: oldProfile.PhoneNumber,
		Gender:      oldProfile.Gender,
		Token:       oldProfile.Token,
	}
	err = service.orchestrator.Start(newProfile, oldAuthProfile)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}

func (service *ProfileService) GetByToken(ctx context.Context, token string) (*domain.Profile, error) {
	return service.store.GetByToken(ctx, token)
}

func (service *ProfileService) GenerateToken(ctx context.Context, profileId string) (string, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return "", err
	}

	return service.store.GenerateToken(ctx, id)
}
