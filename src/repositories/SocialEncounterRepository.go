package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type SocialEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *SocialEncounterRepository) Create(entity *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *SocialEncounterRepository) Delete(id int64) error {
	var entity model.SocialEncounter
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

func (repo *SocialEncounterRepository) Get(id int64) (model.SocialEncounter, error) {
	socialEncounter := model.SocialEncounter{}
	dbResult := repo.DatabaseConnection.First(&socialEncounter, "Id = ?", id)
	if dbResult.Error != nil {
		return socialEncounter, dbResult.Error
	}

	return socialEncounter, nil
}

func (repo *SocialEncounterRepository) GetPaged(page, pageSize int) ([]model.SocialEncounter, error) {
	var entities []model.SocialEncounter
	dbResult := repo.DatabaseConnection.Find(&entities)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return entities, nil
}

func (repo *SocialEncounterRepository) Update(entity *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Save(entity)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}
