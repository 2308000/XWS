package startup

import (
	"accommodation_booking/accommodation_service/application"
	"accommodation_booking/accommodation_service/domain"
	"accommodation_booking/accommodation_service/infrastructure/api"
	"accommodation_booking/accommodation_service/infrastructure/persistence"
	"accommodation_booking/accommodation_service/startup/config"
	"accommodation_booking/common/auth"
	accommodation "accommodation_booking/common/proto/accommodation_service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type Server struct {
	config *config.Config
}

/*const (
	QueueGroup = "accommodation_service"
)*/

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const accommodationServicePath = "/accommodation.AccommodationService/"

	return map[string][]string{
		accommodationServicePath + "GetAll": {"host"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accommodationStore := server.initAccommodationStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	/*commandPublisher := server.initPublisher(server.config.UpdateAccommodationCommandSubject)
	replySubscriber := server.initSubscriber(server.config.UpdateAccommodationReplySubject, QueueGroup)
	updateAccommodationOrchestrator := server.initUpdateAccommodationOrchestrator(commandPublisher, replySubscriber)*/

	accommodationService := server.initAccommodationService(accommodationStore)
	accommodationHandler := server.initAccommodationHandler(accommodationService)

	/*commandSubscriber := server.initSubscriber(server.config.CreateAccommodationCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateAccommodationReplySubject)
	server.initCreateAccommodationHandler(accommodationService, replyPublisher, commandSubscriber)*/

	server.startGrpcServer(accommodationHandler, jwtManager)
}

/*func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}*/

/*func (server *Server) initUpdateAccommodationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UpdateAccommodationOrchestrator {
	orchestrator, err := application.NewUpdateAccommodationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initCreateAccommodationHandler(service *application.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateAccommodationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}*/

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) initAccommodationStore(client *mongo.Client) domain.AccommodationStore {
	store := persistence.NewAccommodationMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, Accommodation := range accommodations {
		acc, err := store.Create(context.TODO(), Accommodation)
		if err != nil {
			log.Fatal(err)
			log.Fatal(acc.Id)
		}
	}
	return store
}

func (server *Server) initAccommodationService(store domain.AccommodationStore) *application.AccommodationService {
	return application.NewAccommodationService(store)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(serverOptions...)
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
