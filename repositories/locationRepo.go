package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"xws_projekat/model"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NoSQL: ProductRepo struct encapsulating Mongo api client
type LocationRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*LocationRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &LocationRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (lr *LocationRepo) Disconnect(ctx context.Context) error {
	err := lr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (lr *LocationRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := lr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		lr.logger.Println(err)
	}

	// Print available databases
	databases, err := lr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		lr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (lr *LocationRepo) GetAll() (model.Locations, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := lr.getCollection()

	var locations model.Locations
	patientsCursor, err := patientsCollection.Find(ctx, bson.M{})
	if err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &locations); err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	return locations, nil
}

func (lr *LocationRepo) GetByCity(city string) (model.Locations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := lr.getCollection()

	var locations model.Locations

	patientsCursor, err := patientsCollection.Find(ctx, bson.M{
		"city": city})

	if err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &locations); err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	return locations, nil
}

func (lr *LocationRepo) GetByCountry(country string) (model.Locations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := lr.getCollection()

	var locations model.Locations

	patientsCursor, err := patientsCollection.Find(ctx, bson.M{
		"country": country})

	if err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &locations); err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	return locations, nil
}

func (lr *LocationRepo) GetByLocation(country, city, airport string) (model.Locations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := lr.getCollection()

	var locations model.Locations

	patientsCursor, err := patientsCollection.Find(ctx, bson.M{
		"city":    city,
		"country": country,
		"airport": airport})

	if err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &locations); err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	return locations, nil
}

func (lr *LocationRepo) Insert(patient *model.Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := lr.getCollection()

	result, err := patientsCollection.InsertOne(ctx, &patient)
	if err != nil {
		lr.logger.Println(err)
		return err
	}
	lr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (lr *LocationRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := lr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := patientsCollection.DeleteOne(ctx, filter)
	if err != nil {
		lr.logger.Println(err)
		return err
	}
	lr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (lr *LocationRepo) getCollection() *mongo.Collection {
	patientDatabase := lr.cli.Database("db")
	patientsCollection := patientDatabase.Collection("patients")
	return patientsCollection
}
