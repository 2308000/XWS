package persistence

import (
	"accommodation_booking/accommodation_service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodation_service"
	COLLECTION = "accommodation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) Get(ctx context.Context, id string) (*domain.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) GetByHost(ctx context.Context, hostId string) ([]*domain.Accommodation, error) {
	id, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		fmt.Println("Host with given Id does not exist!")
	}
	filter := bson.M{"host.hostId": id}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) GetAll(ctx context.Context) ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) Create(ctx context.Context, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	result, err := store.accommodations.InsertOne(ctx, accommodation)
	if err != nil {
		return nil, err
	}
	accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return accommodation, nil
}

func (store *AccommodationMongoDBStore) Update(ctx context.Context, accommodationId string, accommodation *domain.Accommodation) (*domain.Accommodation, error) {
	id, err := primitive.ObjectIDFromHex(accommodationId)
	if err != nil {
		return accommodation, err
	}
	result, err := store.accommodations.ReplaceOne(ctx, bson.M{"_id": id}, accommodation)
	if err != nil {
		return accommodation, err
	}
	if result.MatchedCount == 0 {
		fmt.Println("No document with such ID")
	}
	return accommodation, nil
}

func (store *AccommodationMongoDBStore) Delete(ctx context.Context, accommodationId string) error {
	id, err := primitive.ObjectIDFromHex(accommodationId)
	filter := bson.M{"_id": id}
	_, err = store.accommodations.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *AccommodationMongoDBStore) DeleteAll(ctx context.Context) error {
	_, err := store.accommodations.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
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

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (Accommodation *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&Accommodation)
	return
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var Accommodation domain.Accommodation
		err = cursor.Decode(&Accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &Accommodation)
	}
	err = cursor.Err()
	return
}
