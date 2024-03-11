package main

import (
	"encounters-service/model"
	repository "encounters-service/repositories"
	"fmt"
	"log"
	"time"

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

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	newEncounter, _ := model.NewEncounterExecution(1, 1, 1, time.Now().AddDate(0, 0, -5), time.Now().AddDate(0, 0, -3))
	fmt.Println(newEncounter)
	result := database.Create(&newEncounter)
	if result.Error != nil {
		log.Fatalf("Error creating new encounter: %v", result.Error)
	}

	fmt.Printf("Rows affected: %d\n", result.RowsAffected)
	return database
}

func main() {
	enc := model.Encounter{
		AuthorId: 10,
	}
	fmt.Println(enc)

	database := initDB()
	if database == nil {
		fmt.Println("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repository.EncountersExecutionRepository{DatabaseConnection: database}
	fmt.Println(repo.GetPaged(1, 1))
}
