package main

import (
	"encounters-service/handler"
	"encounters-service/model"
	repository "encounters-service/repositories"
	"encounters-service/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=explorer_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println(database)
	database.AutoMigrate(&model.Encounter{},
		&model.EncounterExecution{},
		&model.EncounterRequest{})

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	return database
}

func startServer(requestHandler *handler.EncounterRequestHandler, encounterHandler *handler.EncounterHandler, executionHandler *handler.EncounterExecutionHandler, touristEncounterHandler *handler.TouristEncounterHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//*requests
	router.HandleFunc("/encounterRequests/getAll", requestHandler.GetAll).Methods("GET")
	router.HandleFunc("/encounterRequests/accept/{id}", requestHandler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/encounterRequests/reject/{id}", requestHandler.RejectRequest).Methods("PUT")

	//*encounters
	router.HandleFunc("/encounter/getAll", encounterHandler.GetAll).Methods("GET")
	router.HandleFunc("/encounter/get/{id}", encounterHandler.GetById).Methods("GET")
	//! update checkpoint
	router.HandleFunc("/encounter/create/{checkpointId}/{isSecretPrerequisite}", encounterHandler.Create).Methods("POST")
	router.HandleFunc("/encounter/update", encounterHandler.Update).Methods("PUT")
	//! update checkpoint
	router.HandleFunc("/encounter/delete/{id}", encounterHandler.Delete).Methods("DELETE")

	//*executions
	router.HandleFunc("/execution/get/{id}", executionHandler.GetById).Methods("GET")
	router.HandleFunc("/execution/getAllByTourist/{id}", executionHandler.GetAllByTourist).Methods("GET")
	router.HandleFunc("/execution/getAllCompletedByTourist/{id}", executionHandler.GetAllCompletedByTourist).Methods("GET")
	//! need body (encounterIds)
	router.HandleFunc("/execution/getByTour/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.GetByTour).Methods("PUT")
	router.HandleFunc("/execution/checkPosition/{id}/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.CheckPosition).Methods("PUT")
	router.HandleFunc("/execution/checkPositionLocationEncounter/{id}/{touristLatitude}/{touristLongitude}/{touristId}", executionHandler.CheckPositionLocationEncounter).Methods("PUT")
	router.HandleFunc("/execution/getActiveByTour/{touristId}", executionHandler.GetActiveByTour).Methods("PUT")
	//! end of body required methods
	router.HandleFunc("/execution/activate/{id}/{touristId}/{touristLatitude}/{touristLongitude}", executionHandler.Activate).Methods("PUT")
	//! update tourists xp points
	router.HandleFunc("/execution/complete/{id}/{touristId}/{touristLatitude}/{touristLongitude}", executionHandler.CompleteExecution).Methods("PUT")
	router.HandleFunc("/execution/update", executionHandler.Update).Methods("PUT")
	router.HandleFunc("/execution/delete/{id}/{touristId}", executionHandler.Update).Methods("DELETE")

	//*tourist encounter
	router.HandleFunc("/touristEncounter/getAll", touristEncounterHandler.GetAll).Methods("GET")
	router.HandleFunc("/touristEncounter/get/{id}", touristEncounterHandler.GetById).Methods("GET")
	//! update checkpoint
	router.HandleFunc("/touristEncounter/create/{checkpointId}/{isSecretPrerequisite}", touristEncounterHandler.Create).Methods("POST")
	router.HandleFunc("/touristEncounter/update", touristEncounterHandler.Update).Methods("PUT")
	//! update checkpoint
	router.HandleFunc("/touristEncounter/delete/{id}", touristEncounterHandler.Delete).Methods("DELETE")

	println("Server listening on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {
	database := initDB()
	if database == nil {
		fmt.Println("FAILED TO CONNECT TO DB")
		return
	}

	encounterRequestRepo := &repository.EncounterRequestRepository{DatabaseConnection: database}
	encounterRequestService := &service.EncounterRequestService{Repo: encounterRequestRepo}
	encounterRequestHandler := &handler.EncounterRequestHandler{EncounterRequestService: encounterRequestService}

	encounterRepo := &repository.EncountersRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{Repo: encounterRepo}
	encounterHandler := &handler.EncounterHandler{Service: encounterService}

	encounterExecutionRepo := &repository.EncountersExecutionRepository{DatabaseConnection: database}
	encounterExecutionService := &service.EncounterExecutionService{Repo: encounterExecutionRepo}
	encounterExecutionHandler := &handler.EncounterExecutionHandler{ExecutionService: encounterExecutionService, EncounterService: encounterService}
	touristEncounterHandler := &handler.TouristEncounterHandler{Service: encounterService, RequestService: encounterRequestService}

	startServer(encounterRequestHandler, encounterHandler, encounterExecutionHandler, touristEncounterHandler)
}
