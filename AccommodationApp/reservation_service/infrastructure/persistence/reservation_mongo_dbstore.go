package persistence

import (
	"accommodation_booking/reservation_service/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "reservation_service"
	COLLECTION = "reservation"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) Get(ctx context.Context, reservationId string) (*domain.Reservation, error) {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) GetAll(ctx context.Context, search string) ([]*domain.Reservation, error) {
	filter := bson.D{{"userId", bson.M{"$regex": "^.*" + search + ".*$"}}}
	return store.filter(filter, search)
}

func (store *ReservationMongoDBStore) Create(ctx context.Context, reservation *domain.Reservation) error {
	result, err := store.reservations.InsertOne(ctx, reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) Update(ctx context.Context, reservationId string, reservation *domain.Reservation) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}
	result, err := store.reservations.ReplaceOne(
		ctx,
		bson.M{"_id": id},
		reservation,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(reservation.Id.String())
	}
	return nil
}

func (store *ReservationMongoDBStore) DeleteAll(ctx context.Context) error {
	_, err := store.reservations.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReservationMongoDBStore) filter(filter interface{}, search string) ([]*domain.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, errors.New(search)
	}
	return decode(cursor)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (reservation *domain.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&reservation)
	return
}

func (store *ReservationMongoDBStore) Delete(ctx context.Context, id string) error {
	reservationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.reservations.DeleteOne(ctx, bson.M{"_id": reservationId})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReservationMongoDBStore) Approve(ctx context.Context, reservationId string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"reservationStatus": domain.ReservationStatusType(1)}}
	_, err = store.reservations.UpdateOne(ctx, filter, update)
	return err
}

func (store *ReservationMongoDBStore) Cancel(ctx context.Context, reservationId string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"reservationStatus": domain.ReservationStatusType(2)}}
	_, err = store.reservations.UpdateOne(ctx, filter, update)
	return err
}

func decode(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var Reservation domain.Reservation
		err = cursor.Decode(&Reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &Reservation)
	}
	err = cursor.Err()
	return
}
