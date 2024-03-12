package main

import (
	"encounters-service/model"
	repository "encounters-service/repositories"
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
	database.AutoMigrate(&model.HiddenLocationEncounter{})
	//database.AutoMigrate(&model.SocialEncounter{})

	err = database.AutoMigrate(&model.Encounter{}, &model.EncounterExecution{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	newEncounter, _ := model.NewEncounter(1, "lara", "cao ja sam lara", 3, 0, 0, 12.321, 33.321)
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
	repo := &repository.EncountersRepository{DatabaseConnection: database}
	addHiddenLocationEncounter(database)
	//addSocialEncounter(database)
	fmt.Println(repo.GetPaged(1, 1))
}

func addHiddenLocationEncounter(database *gorm.DB) {
	newEncounter := model.NewHiddenLocationEncounter(
		model.Encounter{AuthorId: 1, Name: "lara", Description: "cao ja sam lara", Xp: 3, Latitude: 12.321, Longitude: 33.321},
		12.321, 33.321, "example_image.jpg", 10.0,
	)

	result := database.Create(newEncounter)
	if result.Error != nil {
		log.Fatalf("Error creating new HiddenLocationEncounter: %v", result.Error)
	}

	fmt.Printf("New HiddenLocationEncounter ID: %d\n", newEncounter.Id)
}

func addSocialEncounter(database *gorm.DB) {
	newEncounter := model.NewSocialEncounter(
		model.Encounter{AuthorId: 1, Name: "example", Description: "example description", Xp: 3, Latitude: 12.321, Longitude: 33.321},
		5, 10.0,
	)

	result := database.Create(&newEncounter)
	if result.Error != nil {
		log.Fatalf("Error creating new SocialEncounter: %v", result.Error)
	}

	fmt.Printf("New SocialEncounter ID: %d\n", newEncounter.Id)
}
