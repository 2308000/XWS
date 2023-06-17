package persistence

import (
	"accommodation_booking/grade_service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	DATABASE   = "grade_service"
	COLLECTION = "grade"
)

type GradeMongoDBStore struct {
	grades *mongo.Collection
}

func NewGradeMongoDBStore(client *mongo.Client) domain.GradeStore {
	grades := client.Database(DATABASE).Collection(COLLECTION)
	return &GradeMongoDBStore{
		grades: grades,
	}
}

func (store *GradeMongoDBStore) Get(ctx context.Context, gradeId string) (*domain.Grade, error) {
	id, err := primitive.ObjectIDFromHex(gradeId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *GradeMongoDBStore) GetByGuest(ctx context.Context, guestId string) ([]*domain.Grade, error) {
	id, err := primitive.ObjectIDFromHex(guestId)
	if err != nil {
		fmt.Println("Guest with given Id does not exist!")
	}
	filter := bson.M{"guestId": id}
	return store.filter(filter)
}

func (store *GradeMongoDBStore) GetByGraded(ctx context.Context, gradedId string) ([]*domain.Grade, error) {
	id, err := primitive.ObjectIDFromHex(gradedId)
	if err != nil {
		fmt.Println("Host/Accommodation with given Id does not exist!")
	}
	filter := bson.M{"gradedId": id}
	return store.filter(filter)
}

func (store *GradeMongoDBStore) GetAll(ctx context.Context) ([]*domain.Grade, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *GradeMongoDBStore) Create(ctx context.Context, grade *domain.Grade) (*domain.Grade, error) {
	result, err := store.grades.InsertOne(ctx, grade)
	if err != nil {
		return nil, err
	}
	grade.Id = result.InsertedID.(primitive.ObjectID)
	log.Println(grade.Id)
	return grade, nil
}

func (store *GradeMongoDBStore) Update(ctx context.Context, gradeId string, grade *domain.Grade) (*domain.Grade, error) {
	id, err := primitive.ObjectIDFromHex(gradeId)
	if err != nil {
		return grade, err
	}
	result, err := store.grades.ReplaceOne(ctx, bson.M{"_id": id}, grade)
	if err != nil {
		return grade, err
	}
	if result.MatchedCount == 0 {
		fmt.Println("No document with such ID")
	}
	return grade, nil
}

func (store *GradeMongoDBStore) Delete(ctx context.Context, gradeId string) error {
	id, err := primitive.ObjectIDFromHex(gradeId)
	filter := bson.M{"_id": id}
	_, err = store.grades.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *GradeMongoDBStore) DeleteAll(ctx context.Context) error {
	_, err := store.grades.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *GradeMongoDBStore) filter(filter interface{}) ([]*domain.Grade, error) {
	cursor, err := store.grades.Find(context.TODO(), filter)
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

func (store *GradeMongoDBStore) filterOne(filter interface{}) (Grade *domain.Grade, err error) {
	result := store.grades.FindOne(context.TODO(), filter)
	err = result.Decode(&Grade)
	return
}

func decode(cursor *mongo.Cursor) (grades []*domain.Grade, err error) {
	for cursor.Next(context.TODO()) {
		var Grade domain.Grade
		err = cursor.Decode(&Grade)
		if err != nil {
			return
		}
		grades = append(grades, &Grade)
	}
	err = cursor.Err()
	return
}
