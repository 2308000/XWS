package client

import (
	pbAccommodation "accommodation_booking/common/proto/accommodation_service"
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
