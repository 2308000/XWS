package persistence

import (
	auth "accommodation_booking/common/domain"
	"accommodation_booking/user_service/domain"
	"accommodation_booking/user_service/infrastructure/api"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "user_service"
	COLLECTION1 = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION1)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(ctx context.Context, username string) (*auth.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll(ctx context.Context) ([]*auth.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Register(ctx context.Context, user *auth.User) (*auth.User, error) {
	result, err := store.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *UserMongoDBStore) Update(ctx context.Context, id primitive.ObjectID, username string) (string, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": username}}
	_, err := store.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (store *UserMongoDBStore) Delete(ctx context.Context, id primitive.ObjectID) error {

	filter := bson.M{"_id": id}
	_, err := store.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) DeleteAll(ctx context.Context) error {
	_, err := store.users.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*auth.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *auth.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (users []*auth.User, err error) {
	for cursor.Next(context.TODO()) {
		var User auth.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}

func (store *UserMongoDBStore) UpdatePassword(ctx context.Context, username string, password string) (string, error) {
	_, err := store.users.UpdateOne(ctx, bson.M{"username": username},
		bson.M{"$set": bson.M{"password": api.HashPassword(password)}})
	if err != nil {
		return "", err
	}
	return "Password has been changed successfully", nil
}
