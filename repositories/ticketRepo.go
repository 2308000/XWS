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
type TicketRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewTicketRepo(ctx context.Context, logger *log.Logger) (*TicketRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &TicketRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (tr *TicketRepo) Disconnect(ctx context.Context) error {
	err := tr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (tr *TicketRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := tr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		tr.logger.Println(err)
	}

	// Print available databases
	databases, err := tr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		tr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (tr *TicketRepo) GetAll() ([]*model.Ticket, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := tr.getCollection()

	var tickets []*model.Ticket
	ticketsCursor, err := ticketsCollection.Find(ctx, bson.M{})
	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	if err = ticketsCursor.All(ctx, &tickets); err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return tickets, nil
}

func (tr *TicketRepo) GetById(id string) (*model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := tr.getCollection()

	var ticket model.Ticket
	objID, _ := primitive.ObjectIDFromHex(id)
	err := ticketsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ticket)
	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return &ticket, nil
}

func (tr *TicketRepo) GetByUser(user *model.User) ([]*model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := tr.getCollection()

	var tickets []*model.Ticket
	ticketsCursor, err := ticketsCollection.Find(ctx, bson.M{"user": user})
	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	if err = ticketsCursor.All(ctx, &tickets); err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return tickets, nil
}

func (tr *TicketRepo) Insert(ticket *model.Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	result, err := ticketsCollection.InsertOne(ctx, &ticket)
	if err != nil {
		tr.logger.Println(err)
		return err
	}
	tr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (tr *TicketRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := ticketsCollection.DeleteOne(ctx, filter)
	if err != nil {
		tr.logger.Println(err)
		return err
	}
	tr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (tr *TicketRepo) getCollection() *mongo.Collection {
	ticketDatabase := tr.cli.Database("db")
	ticketsCollection := ticketDatabase.Collection("tickets")
	return ticketsCollection
}
