package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type EncounterRequestRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRequestRepository) Create(encounterRequest *model.EncounterRequest) error {
	dbResult := repo.DatabaseConnection.Create(encounterRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncounterRequestRepository) AcceptRequest(id int64) (model.EncounterRequest, error) {
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

func (repo *EncounterRequestRepository) RejectRequest(id int64) (model.EncounterRequest, error) {
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

func (repo *EncounterRequestRepository) GetAll() ([]model.EncounterRequest, error) {
	var encounterRequests []model.EncounterRequest
	dbResult := repo.DatabaseConnection.Find(&encounterRequests)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounterRequests, nil
}
