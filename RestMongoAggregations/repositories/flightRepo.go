package repositories

import (
	"Rest/model"
	"context"
	"fmt"
	"log"
	"time"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NoSQL: ProductRepo struct encapsulating Mongo api client
type FlightRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewFlightRepo(ctx context.Context, logger *log.Logger) (*FlightRepo, error) {
	//dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:pass@mongo:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &FlightRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (fr *FlightRepo) Disconnect(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (fr *FlightRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := fr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		fr.logger.Println(err)
	}

	// Print available databases
	databases, err := fr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (fr *FlightRepo) GetAll() (model.Flights, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	var flights model.Flights

	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepo) GetFlightsByDate(flightDate time.Time) (model.Flights, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dateFilterStart := time.Date(flightDate.Year(), flightDate.Month(), flightDate.Day(), 0, 0, 0, 0, time.Local)
	dateFilterEnd := time.Date(flightDate.Year(), flightDate.Month(), flightDate.Day(), 23, 59, 99, 0, time.Local)

	flightsCollection := fr.getCollection()

	var flights model.Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{"startTime": bson.M{
		"$gte": dateFilterStart.Format(time.RFC3339),
		"$lte": dateFilterEnd.Format(time.RFC3339),
	}})

	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepo) GetById(id string) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	var flight model.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)

	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return &flight, nil
}

func (fr *FlightRepo) Insert(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (fr *FlightRepo) Update(id string, flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"sold": flight.TicketsSold,
	}}
	result, err := flightsCollection.UpdateOne(ctx, filter, update)
	fr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	fr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		fr.logger.Println(err)
		return err
	}
	return nil
}

func (fr *FlightRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (fr *FlightRepo) getCollection() *mongo.Collection {
	flightDatabase := fr.cli.Database("mongoDemo")
	flightsCollection := flightDatabase.Collection("flights")

	return flightsCollection
}
