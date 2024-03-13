package main

import (
	"encounters-service/handler"
	"encounters-service/model"
	repository "encounters-service/repositories"
	"encounters-service/service"
	"fmt"
	"log"

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

	router.HandleFunc("/encounters/getAll", handler.GetAll).Methods("GET")

	println("Server starting")
	//log.Fatal(http.ListenAndServe(":BIRACEMO_PORT", router))
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
