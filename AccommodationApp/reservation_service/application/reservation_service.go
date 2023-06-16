package application

import (
	//auth "accommodation_booking/common/domain"
	"accommodation_booking/reservation_service/domain"
	"context"
	"errors"
	"time"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(ctx context.Context, reservationId string) (*domain.Reservation, error) {
	return service.store.Get(ctx, reservationId)
}

func (service *ReservationService) GetAll(ctx context.Context) ([]*domain.Reservation, error) {
	return service.store.GetAll(ctx)
}

func (service *ReservationService) GetForUser(ctx context.Context, userId string) ([]*domain.Reservation, error) {
	return service.store.GetForUser(ctx, userId)
}

func (service *ReservationService) Create(ctx context.Context, reservation *domain.Reservation) error {
	reservations, err := service.store.GetBetweenDates(ctx, reservation.Beginning, reservation.Ending, reservation.AccommodationId.Hex())
	if err != nil {
		return err
	}
	if len(reservations) != 0 {
		return errors.New("there is reservation overlapping selected time interval")
	}
	return service.store.Create(ctx, reservation)
}

func (service *ReservationService) RollbackUpdate(ctx context.Context, reservation *domain.Reservation) error {
	return service.store.Update(ctx, reservation.Id.Hex(), reservation)
}

func (service *ReservationService) Update(ctx context.Context, reservationId string, reservation *domain.Reservation) error {
	err := service.store.Update(ctx, reservationId, reservation)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Approve(ctx context.Context, reservationId string) error {
	err := service.store.Approve(ctx, reservationId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Cancel(ctx context.Context, reservationId string) error {
	reservation, err := service.store.Get(ctx, reservationId)
	if err != nil {
		return err
	}
	if reservation.UserId.Hex() != ctx.Value("userId").(string) {
		return errors.New("you are not allowed cancel this reservation")
	}
	tomorrow := time.Now().Add(1)
	if reservation.Beginning.Before(tomorrow) {
		return errors.New("you cannot cancel a reservation if there is less than a day left until it starts")
	}
	if reservation.ReservationStatus == 0 {
		err = service.Delete(ctx, reservationId)
	} else {
		err = service.store.Cancel(ctx, reservationId)
	}
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) Delete(ctx context.Context, id string) error {
	return service.store.Delete(ctx, id)
}
