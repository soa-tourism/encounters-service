package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type EncounterRequestDatabaseRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRequestDatabaseRepository) Create(encounterRequest *model.EncounterRequest) error {
	dbResult := repo.DatabaseConnection.Create(encounterRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncounterRequestDatabaseRepository) AcceptRequest(id int64) (model.EncounterRequest, error) {
	requestToUpdate := model.EncounterRequest{}
	dbResult := repo.DatabaseConnection.First(&requestToUpdate, "Id = ?", id)
	if dbResult.Error != nil {
		return requestToUpdate, dbResult.Error
	}

	requestToUpdate.AcceptRequest()
	dbResult2 := repo.DatabaseConnection.Save(&requestToUpdate)
	if dbResult2.Error != nil {
		return requestToUpdate, dbResult2.Error
	}

	return requestToUpdate, nil
}

func (repo *EncounterRequestDatabaseRepository) RejectRequest(id int64) (model.EncounterRequest, error) {
	requestToUpdate := model.EncounterRequest{}
	dbResult := repo.DatabaseConnection.First(&requestToUpdate, "Id = ?", id)
	if dbResult.Error != nil {
		return requestToUpdate, dbResult.Error
	}

	requestToUpdate.RejectRequest()
	dbResult2 := repo.DatabaseConnection.Save(&requestToUpdate)
	if dbResult2.Error != nil {
		return requestToUpdate, dbResult2.Error
	}

	return requestToUpdate, nil
}
