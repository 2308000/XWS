package handlers

import (
	//"Rest/data"
	"context"
	"log"
	"net/http"
	"xws_projekat/model"
	"xws_projekat/repositories"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type FlightsHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	flightRepo   *repositories.FlightRepo
	locationRepo *repositories.LocationRepo
}

// Injecting the logger makes this code much more testable.
func NewFlightsHandler(l *log.Logger, fr *repositories.FlightRepo, lr *repositories.LocationRepo) *FlightsHandler {
	return &FlightsHandler{l, fr, lr}
}

func (f *FlightsHandler) GetAllFlights(rw http.ResponseWriter, h *http.Request) {
	flights, err := f.flightRepo.GetAll()
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (f *FlightsHandler) CreateFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)
	f.flightRepo.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (p *FlightsHandler) UpdateFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)

	p.flightRepo.Update(id, flight)
	rw.WriteHeader(http.StatusOK)
}

func (p *FlightsHandler) DeleteFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.flightRepo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (f *FlightsHandler) MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (f *FlightsHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		f.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
