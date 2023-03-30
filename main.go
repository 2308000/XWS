package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"xws_projekat/handlers"
	"xws_projekat/repositories"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[flight-store] ", log.LstdFlags)

	store, err := repositories.NewFlightRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	store.Ping()

	//Initialize the handler and inject said logger
	flightsHandler := handlers.NewFlightsHandler(logger, store)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(flightsHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", flightsHandler.GetAllFlights)
	/*flight1 := model.Flight{primitive.NewObjectID(), "Madrid", "Belgrade", time.Now(), 100.0, 150, 120}
	store.Insert(&flight1)
	flight2 := model.Flightprimitive.NewObjectID(), "Belgrade", "Vienna", time.Now(), 100.0, 120, 110)
	flight3 := model.Flight(primitive.NewObjectID(), "Vienna", "Berlin", time.Now(), 100.0, 110, 100)*/

	getFlightRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightRouter.HandleFunc("/", flightsHandler.GetAllFlights)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", flightsHandler.CreateFlight)
	postRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	/*getByNameRouter := router.Methods(http.MethodGet).Subrouter()
	getByNameRouter.HandleFunc("/filter", flightsHandler.GetPatientsByName)

	receiptRouter := router.Methods(http.MethodGet).Subrouter()
	receiptRouter.HandleFunc("/receipt/{id}", flightsHandler.Receipt)

	reportRouter := router.Methods(http.MethodGet).Subrouter()
	reportRouter.HandleFunc("/report", flightsHandler.Report)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", flightsHandler.GetPatientById)

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/{id}", flightsHandler.PatchPatient)
	patchRouter.Use(flightsHandler.MiddlewarePatientDeserialization)

	changePhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	changePhoneRouter.HandleFunc("/phone/{id}/{index}", flightsHandler.ChangePhone)

	pushPhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	pushPhoneRouter.HandleFunc("/phone/{id}", flightsHandler.AddPhoneNumber)

	addAnamnesisRouter := router.Methods(http.MethodPatch).Subrouter()
	addAnamnesisRouter.HandleFunc("/anamnesis/{id}", flightsHandler.AddAnamnesis)

	addTherapyRouter := router.Methods(http.MethodPatch).Subrouter()
	addTherapyRouter.HandleFunc("/therapy/{id}", flightsHandler.AddTherapy)

	changeAddressRouter := router.Methods(http.MethodPatch).Subrouter()
	changeAddressRouter.HandleFunc("/address/{id}", flightsHandler.ChangeAddress)*/
	getByNameRouter.HandleFunc("/filter", patientsHandler.GetPatientsByName)

	receiptRouter := router.Methods(http.MethodGet).Subrouter()
	receiptRouter.HandleFunc("/receipt/{id}", patientsHandler.Receipt)

	reportRouter := router.Methods(http.MethodGet).Subrouter()
	reportRouter.HandleFunc("/report", patientsHandler.Report)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", patientsHandler.GetPatientById)*/

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/{id}", flightsHandler.UpdateFlight)
	postRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	/*changePhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	changePhoneRouter.HandleFunc("/phone/{id}/{index}", patientsHandler.ChangePhone)

	pushPhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	pushPhoneRouter.HandleFunc("/phone/{id}", patientsHandler.AddPhoneNumber)

	addAnamnesisRouter := router.Methods(http.MethodPatch).Subrouter()
	addAnamnesisRouter.HandleFunc("/anamnesis/{id}", patientsHandler.AddAnamnesis)

	addTherapyRouter := router.Methods(http.MethodPatch).Subrouter()
	addTherapyRouter.HandleFunc("/therapy/{id}", patientsHandler.AddTherapy)

	changeAddressRouter := router.Methods(http.MethodPatch).Subrouter()
	changeAddressRouter.HandleFunc("/address/{id}", patientsHandler.ChangeAddress)*/

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", flightsHandler.DeleteFlight)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
