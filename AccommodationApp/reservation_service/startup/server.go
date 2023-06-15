package startup

import (
	"accommodation_booking/common/auth"
	reservation "accommodation_booking/common/proto/reservation_service"
	"accommodation_booking/reservation_service/application"
	"accommodation_booking/reservation_service/domain"
	"accommodation_booking/reservation_service/infrastructure/api"
	"accommodation_booking/reservation_service/infrastructure/persistence"
	"accommodation_booking/reservation_service/startup/config"
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

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const reservationServicePath = "/reservation.ReservationService/"

	return map[string][]string{
		reservationServicePath + "GetAll":  {"guest", "host"},
		reservationServicePath + "Get":     {"guest", "host"},
		reservationServicePath + "Approve": {"host"},
		reservationServicePath + "Cancel":  {"guest"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationStore := server.initReservationStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	reservationService := server.initReservationService(reservationStore)
	reservationHandler := server.initReservationHandler(reservationService)

	server.startGrpcServer(reservationHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}

func (server *Server) initReservationStore(client *mongo.Client) domain.ReservationStore {
	store := persistence.NewReservationMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, Reservation := range reservations {
		err := store.Create(context.TODO(), Reservation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initReservationService(store domain.ReservationStore) *application.ReservationService {
	return application.NewReservationService(store)
}

func (server *Server) startGrpcServer(reservationHandler *api.ReservationHandler, jwtManager *auth.JWTManager) {
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
	reservation.RegisterReservationServiceServer(grpcServer, reservationHandler)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
