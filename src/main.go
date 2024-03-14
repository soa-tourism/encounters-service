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
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=encounters"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	database.AutoMigrate(&model.Encounter{},
		&model.EncounterExecution{},
		&model.EncounterRequest{})

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startEncounterServer(handler *handler.EncounterRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounterRequests/getAll", handler.GetAll).Methods("GET")
	router.HandleFunc("/encounterRequests/accept/{id}", handler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/encounterRequests/reject/{id}", handler.RejectRequest).Methods("PUT")

	println("Server starting")
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
	startEncounterServer(encounterRequestHandler)
}
