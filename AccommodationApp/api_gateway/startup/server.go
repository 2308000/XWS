package startup

import (
	cfg "accommodation_booking/api_gateway/startup/config"
	accommodationGw "accommodation_booking/common/proto/accommodation_service"
	profileGw "accommodation_booking/common/proto/profile_service"
	reservationGw "accommodation_booking/common/proto/reservation_service"
	userGw "accommodation_booking/common/proto/user_service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	profileEndpoint := fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort)
	err = profileGw.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, profileEndpoint, opts)
	if err != nil {
		panic(err)
	}
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err = reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(server.mux)))

}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		if r.Header.Get("Origin") != "" {
			h.Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		//h.Set("Access-Control-Allow-Origin", "https://localhost:7777")

		if r.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodOptions,
					http.MethodGet,
					http.MethodPut,
					http.MethodHead,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodTrace,
				}, ", ",
			))

			h.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
					"Authorization",
					"Location",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}
