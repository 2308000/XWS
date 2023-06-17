package startup

import (
	"accommodation_booking/accommodation_service/application"
	"accommodation_booking/accommodation_service/domain"
	"accommodation_booking/accommodation_service/infrastructure/api"
	"accommodation_booking/accommodation_service/infrastructure/persistence"
	"accommodation_booking/accommodation_service/startup/config"
	"accommodation_booking/common/auth"
	"accommodation_booking/common/client"
	accommodation "accommodation_booking/common/proto/accommodation_service"
	grade "accommodation_booking/common/proto/grade_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	saga "accommodation_booking/common/saga/messaging"
	"accommodation_booking/common/saga/messaging/nats"
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

const (
	QueueGroup = "accommodation_service"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const accommodationServicePath = "/accommodation.AccommodationService/"

	return map[string][]string{
		accommodationServicePath + "GetAll":             {"host"},
		accommodationServicePath + "Create":             {"host"},
		accommodationServicePath + "UpdateAvailability": {"host"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accommodationStore := server.initAccommodationStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	reservationClient, err := client.NewReservationClient(fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	gradeClient, err := client.NewGradeClient(fmt.Sprintf("%s:%s", server.config.GradeHost, server.config.GradePort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	userClient, err := client.NewUserClient(fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	accommodationService := server.initAccommodationService(accommodationStore)
	accommodationHandler := server.initAccommodationHandler(accommodationService, profileClient, reservationClient, gradeClient, userClient)

	commandSubscriber := server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(accommodationService, replyPublisher, commandSubscriber)

	server.startGrpcServer(accommodationHandler, jwtManager)
}

func (server *Server) initUpdateProfileHandler(service *application.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
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
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationHandler(service *application.AccommodationService, profileClient profile.ProfileServiceClient, reservationClient reservation.ReservationServiceClient, gradeClient grade.GradeServiceClient, userClient user.UserServiceClient) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service, profileClient, reservationClient, gradeClient, userClient)
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
