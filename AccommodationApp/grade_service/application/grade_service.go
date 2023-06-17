package application

import (
	"accommodation_booking/grade_service/domain"
	"context"
)

type GradeService struct {
	store domain.GradeStore
}

func NewGradeService(store domain.GradeStore) *GradeService {
	return &GradeService{
		store: store,
	}
}

func (service *GradeService) Get(ctx context.Context, gradeId string) (*domain.Grade, error) {
	return service.store.Get(ctx, gradeId)
}

func (service *GradeService) GetByGuest(ctx context.Context, guestId string) ([]*domain.Grade, error) {
	return service.store.GetByGuest(ctx, guestId)
}

func (service *GradeService) GetByGraded(ctx context.Context, gradedId string) ([]*domain.Grade, error) {
	return service.store.GetByGuest(ctx, gradedId)
}

func (service *GradeService) GetAll(ctx context.Context) ([]*domain.Grade, error) {
	return service.store.GetAll(ctx)
}

func (service *GradeService) Create(ctx context.Context, grade *domain.Grade) (*domain.Grade, error) {
	return service.store.Create(ctx, grade)
}

func (service *GradeService) Update(ctx context.Context, acommodationId string, grade *domain.Grade) (*domain.Grade, error) {
	return service.store.Update(ctx, acommodationId, grade)
}

func (service *GradeService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}
