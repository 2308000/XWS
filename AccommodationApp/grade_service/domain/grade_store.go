package domain

import (
	"context"
)

type GradeStore interface {
	Get(ctx context.Context, id string) (*Grade, error)
	GetHostsGradedByGuest(ctx context.Context, guestId string) ([]*Grade, error)
	GetAccommodationsGradedByGuest(ctx context.Context, guestId string) ([]*Grade, error)
	GetByGraded(ctx context.Context, gradedId string) ([]*Grade, error)
	GetAll(ctx context.Context) ([]*Grade, error)
	Create(ctx context.Context, accommodation *Grade) (*Grade, error)
	Update(ctx context.Context, accommodationId string, accommodation *Grade) (*Grade, error)
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error
}
