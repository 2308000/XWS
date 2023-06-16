package application

import (
	auth "accommodation_booking/common/domain"
	"accommodation_booking/user_service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store        domain.UserStore
	orchestrator *CreateProfileOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *CreateProfileOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *UserService) Get(ctx context.Context, username string) (*auth.User, error) {
	return service.store.Get(ctx, username)
}

func (service *UserService) GetById(ctx context.Context, id string) (*auth.User, error) {
	return service.store.GetById(ctx, id)
}

func (service *UserService) GetAll(ctx context.Context) ([]*auth.User, error) {
	return service.store.GetAll(ctx)
}

func (service *UserService) Register(ctx context.Context, user *auth.User, firstName string, lastName string, email string, address auth.Address) (*auth.User, error) {
	registeredUser, err := service.store.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	profile := &auth.Profile{
		Id:        registeredUser.Id,
		Username:  registeredUser.Username,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  firstName + lastName,
		Email:     email,
		Address: auth.Address{
			Country: address.Country,
			City:    address.City,
			Street:  address.Street,
		},
	}
	service.orchestrator.Start(profile)
	return registeredUser, nil
}

func (service *UserService) Update(ctx context.Context, id primitive.ObjectID, username string) (string, error) {
	return service.store.Update(ctx, id, username)
}

func (service *UserService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return service.store.Delete(ctx, id)
}

func (service *UserService) UpdatePassword(ctx context.Context, username string, password string) error {
	_, err := service.store.UpdatePassword(ctx, username, password)
	if err != nil {
		return err
	}
	return nil
}
