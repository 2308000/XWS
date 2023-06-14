package application

import (
	"accommodation_booking/accommodation_service/domain"
	"context"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(ctx context.Context, accommodationId string) (*domain.Accommodation, error) {
	return service.store.Get(ctx, accommodationId)
}

func (service *AccommodationService) GetByHost(ctx context.Context, hostId string) ([]*domain.Accommodation, error) {
	return service.store.GetByHost(ctx, hostId)
}

func (service *AccommodationService) GetAll(ctx context.Context) ([]*domain.Accommodation, error) {
	return service.store.GetAll(ctx)
}

func (service *AccommodationService) Create(ctx context.Context, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	return service.store.Create(ctx, accommodation)
}

func (service *AccommodationService) Update(ctx context.Context, acommodationId string, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	return service.store.Update(ctx, acommodationId, accommodation)
}

func (service *AccommodationService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}
