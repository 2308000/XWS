package domain

import (
	"context"
)

type AccommodationStore interface {
	Get(ctx context.Context, id string) (*Accommodation, error)
	GetByHost(ctx context.Context, hostId string) ([]*Accommodation, error)
	GetAll(ctx context.Context) ([]*Accommodation, error)
	Create(ctx context.Context, accommodation *Accommodation) (*Accommodation, error)
	Update(ctx context.Context, accommodationId string, accommodation *Accommodation) (*Accommodation, error)
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error
}
