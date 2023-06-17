package startup

import (
	"accommodation_booking/common/auth"
	"accommodation_booking/common/client"
	grade "accommodation_booking/common/proto/grade_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	user "accommodation_booking/common/proto/user_service"
	saga "accommodation_booking/common/saga/messaging"
	"accommodation_booking/common/saga/messaging/nats"
	"accommodation_booking/profile_service/application"
	"accommodation_booking/profile_service/domain"
	"accommodation_booking/profile_service/infrastructure/api"
	"accommodation_booking/profile_service/infrastructure/persistence"
	"accommodation_booking/profile_service/startup/config"
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
	QueueGroup = "profile_service"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const profileServicePath = "/profile.ProfileService/"

	return map[string][]string{
		profileServicePath + "GetAll": {"guest", "host"},
		profileServicePath + "Delete": {"guest", "host"},
		profileServicePath + "Update": {"guest", "host"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	profileStore := server.initProfileStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	commandPublisher := server.initPublisher(server.config.UpdateProfileCommandSubject)
	replySubscriber := server.initSubscriber(server.config.UpdateProfileReplySubject, QueueGroup)
	updateProfileOrchestrator := server.initUpdateProfileOrchestrator(commandPublisher, replySubscriber)

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

	profileService := server.initProfileService(profileStore, updateProfileOrchestrator)
	profileHandler := server.initProfileHandler(profileService, reservationClient, gradeClient, userClient)

	commandSubscriber := server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateProfileReplySubject)
	server.initCreateProfileHandler(profileService, replyPublisher, commandSubscriber)

	server.startGrpcServer(profileHandler, jwtManager)
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

func (server *Server) initUpdateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UpdateProfileOrchestrator {
	orchestrator, err := application.NewUpdateProfileOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initCreateProfileHandler(service *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProfileHandler(service *application.ProfileService, reservationClient reservation.ReservationServiceClient, gradeClient grade.GradeServiceClient, userClient user.UserServiceClient) *api.ProfileHandler {
	return api.NewProfileHandler(service, reservationClient, gradeClient, userClient)
}

func (server *Server) initProfileStore(client *mongo.Client) domain.ProfileStore {
	store := persistence.NewProfileMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, Profile := range profiles {
		err := store.Create(context.TODO(), Profile)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initProfileService(store domain.ProfileStore, orchestrator *application.UpdateProfileOrchestrator) *application.ProfileService {
	return application.NewProfileService(store, orchestrator)
}

func (server *Server) startGrpcServer(profileHandler *api.ProfileHandler, jwtManager *auth.JWTManager) {
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
	profile.RegisterProfileServiceServer(grpcServer, profileHandler)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
