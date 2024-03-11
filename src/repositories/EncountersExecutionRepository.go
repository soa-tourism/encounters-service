package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type EncountersExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncountersExecutionRepository) GetPaged(page, pageSize int) ([]model.EncounterExecution, error) {
	var executions []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Find(&executions)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return executions, nil
}

func (repo *EncountersExecutionRepository) Create(encounterExecution *model.EncounterExecution) error {
	dbResult := repo.DatabaseConnection.Create(encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncountersExecutionRepository) Get(id int64) (model.EncounterExecution, error) {
	encounterExecution := model.EncounterExecution{}
	dbResult := repo.DatabaseConnection.Preload("Encounter").First(&encounterExecution, "Id = ?", id)
	if dbResult.Error != nil {
		return encounterExecution, dbResult.Error
	}

	return encounterExecution, nil
}

func (repo *EncountersExecutionRepository) Update(encounterExecution *model.EncounterExecution) error {
	dbResult := repo.DatabaseConnection.Save(encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncountersExecutionRepository) Delete(id int64) error {
	var encounterExecution model.EncounterExecution
	dbResult := repo.DatabaseConnection.First(&encounterExecution, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	dbResult = repo.DatabaseConnection.Delete(&encounterExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncountersExecutionRepository) GetAllByTourist(touristId int64) ([]model.EncounterExecution, error) {
	var encounters []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("TouristId = ?", touristId).Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) GetAllCompletedByTourist(touristId int64) ([]model.EncounterExecution, error) {
	var encounters []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("TouristId = ? AND Status = 1", touristId).Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) FindByEncounterId(encounterId int64) (model.EncounterExecution, error) {
	var execution model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("EncounterId = ?", encounterId).First(&execution)

	if dbResult.Error != nil {
		return execution, dbResult.Error
	}

	return execution, nil
}

func (repo *EncountersExecutionRepository) GetByEncounterAndTourist(touristId, encounterId int64) (model.EncounterExecution, error) {
	var execution model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("EncounterId = ? AND TouristId = ? ", encounterId, touristId).First(&execution)

	if dbResult.Error != nil {
		return execution, dbResult.Error
	}

	return execution, nil
}

func (repo *EncountersExecutionRepository) GetActiveByTourist(touristId int64) ([]model.EncounterExecution, error) {
	var encounters []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("TouristId = ? AND Status = 2", touristId).Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) UpdateRange(encounters []model.EncounterExecution) ([]model.EncounterExecution, error) {
	tx := repo.DatabaseConnection.Begin()

	for _, encounter := range encounters {
		if err := tx.Save(&encounter).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) GetBySocialEncounter(socialEncounterId int64) ([]model.EncounterExecution, error) {
	var encounters []model.EncounterExecution
	dbResult := repo.DatabaseConnection.
		Preload("Encounter", "Type = 0").
		Where("EncounterId = ?", socialEncounterId).
		Find(&encounters)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) GetByLocationEncounter(locationEncounterId int64) ([]model.EncounterExecution, error) {
	var encounters []model.EncounterExecution
	dbResult := repo.DatabaseConnection.
		Preload("Encounter", "Type = 1").
		Where("EncounterId = ?", locationEncounterId).
		Find(&encounters)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersExecutionRepository) GetByEncounter(encounterId int64) ([]model.EncounterExecution, error) {
	var executions []model.EncounterExecution
	dbResult := repo.DatabaseConnection.Preload("Encounter").Where("EncounterId = ?", encounterId).Find(&executions)
	if dbResult.Error != nil {
		return executions, dbResult.Error
	}
	return executions, nil
}
