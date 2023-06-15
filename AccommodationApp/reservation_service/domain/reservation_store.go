package domain

import "context"

type ReservationStore interface {
	Get(ctx context.Context, reservationId string) (*Reservation, error)
	GetAll(ctx context.Context, search string) ([]*Reservation, error)
	Create(ctx context.Context, reservation *Reservation) error
	Update(ctx context.Context, reservationId string, reservation *Reservation) error
	Approve(ctx context.Context, reservationId string) error
	DeleteAll(ctx context.Context) error
	Delete(ctx context.Context, id string) error
}
