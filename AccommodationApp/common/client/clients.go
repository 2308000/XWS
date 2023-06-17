package client

import (
	pbAccommodation "accommodation_booking/common/proto/accommodation_service"
	pbGrade "accommodation_booking/common/proto/grade_service"
	pbProfile "accommodation_booking/common/proto/profile_service"
	pbReservation "accommodation_booking/common/proto/reservation_service"
	pbUser "accommodation_booking/common/proto/user_service"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"google.golang.org/grpc"
)

func NewUserClient(address string) (pbUser.UserServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbUser.NewUserServiceClient(conn)
	return client, nil
}

func NewAccommodationClient(address string) (pbAccommodation.AccommodationServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbAccommodation.NewAccommodationServiceClient(conn)
	return client, nil
}

func NewProfileClient(address string) (pbProfile.ProfileServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbProfile.NewProfileServiceClient(conn)
	return client, nil
}

func NewReservationClient(address string) (pbReservation.ReservationServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbReservation.NewReservationServiceClient(conn)
	return client, nil
}

func NewGradeClient(address string) (pbGrade.GradeServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbGrade.NewGradeServiceClient(conn)
	return client, nil
}
