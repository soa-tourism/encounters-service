package main

import (
	"encounters-service/dto"
	"encounters-service/model"
	repository "encounters-service/repositories"
	"encounters-service/service"
	"fmt"
	"log"

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

	database.AutoMigrate(&model.Encounter{})
	database.AutoMigrate(&model.EncounterExecution{})
	database.AutoMigrate(&model.EncounterRequest{})

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	// result := database.Create(&newEncounter)
	// if result.Error != nil {
	// 	log.Fatalf("Error creating new encounter: %v", result.Error)
	// }
	// fmt.Printf("Rows affected: %d\n", result.RowsAffected)
	return database
}

func main() {
	database := initDB()
	if database == nil {
		fmt.Println("FAILED TO CONNECT TO DB")
		return
	}
	newEncounter := dto.EncounterDto{
		Id:                0,
		AuthorId:          1,
		Name:              "Enc",
		Description:       "Enco",
		Xp:                12,
		Status:            0,
		Type:              0,
		Longitude:         12.2,
		Latitude:          15.5,
		LocationLongitude: 0,
		LocationLatitude:  0,
		Image:             "",
		Range:             10.0,
		ActiveTouristsIds: make([]int, 0),
		RequiredPeople:    4,
	}

	repo := &repository.EncountersRepository{DatabaseConnection: database}
	service := &service.EncounterService{
		Repo: repo,
	}
	encounter, _ := service.Create(newEncounter, 10, true, 1)
	fmt.Println(encounter)
	fmt.Println(repo.GetPaged(1, 1))
}
