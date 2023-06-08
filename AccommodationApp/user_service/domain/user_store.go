package domain

import (
	auth "accommodation_booking/common/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(ctx context.Context, username string) (*auth.User, error)
	GetAll(ctx context.Context) ([]*auth.User, error)
	Register(ctx context.Context, user *auth.User) (*auth.User, error)
	Update(ctx context.Context, id primitive.ObjectID, username string) (string, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	DeleteAll(ctx context.Context) error
	UpdatePassword(ctx context.Context, username string, password string) (string, error)
}
