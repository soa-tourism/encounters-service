package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type HiddenLocationEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *HiddenLocationEncounterRepository) Create(entity *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *HiddenLocationEncounterRepository) Delete(id int64) error {
	var entity model.HiddenLocationEncounter
	dbResult := repo.DatabaseConnection.First(&entity, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	dbResult = repo.DatabaseConnection.Delete(&entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *HiddenLocationEncounterRepository) Get(id int64) (model.HiddenLocationEncounter, error) {
	hiddenLocationEncounter := model.HiddenLocationEncounter{}
	dbResult := repo.DatabaseConnection.First(&hiddenLocationEncounter, "Id = ?", id)
	if dbResult.Error != nil {
		return hiddenLocationEncounter, dbResult.Error
	}

	return hiddenLocationEncounter, nil
}

func (repo *HiddenLocationEncounterRepository) GetPaged(page, pageSize int) ([]model.HiddenLocationEncounter, error) {
	var entities []model.HiddenLocationEncounter
	dbResult := repo.DatabaseConnection.Find(&entities)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return entities, nil
}

func (repo *HiddenLocationEncounterRepository) Update(entity *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Save(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
