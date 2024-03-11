package repository

import (
	"encounters-service/model"

	"gorm.io/gorm"
)

type EncountersRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncountersRepository) Create(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncountersRepository) Delete(id int64) error {
	var encounter model.Encounter
	dbResult := repo.DatabaseConnection.First(&encounter, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	dbResult = repo.DatabaseConnection.Delete(&encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncountersRepository) Get(id int64) (model.Encounter, error) {
	encounter := model.Encounter{}
	dbResult := repo.DatabaseConnection.First(&encounter, "Id = ?", id)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *EncountersRepository) GetPaged(page, pageSize int) ([]model.Encounter, error) {
	var encounters []model.Encounter
	dbResult := repo.DatabaseConnection.Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}

func (repo *EncountersRepository) MakeEncounterPublished(encounterId int) (model.Encounter, error) {
	encounter := model.Encounter{}
	dbResult := repo.DatabaseConnection.First(&encounter, "Id = ?", encounterId)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}

	encounter.MakeEncounterPublished()
	dbResult2 := repo.DatabaseConnection.Save(encounter)
	if dbResult2.Error != nil {
		return encounter, dbResult2.Error
	}

	return encounter, nil
}

func (repo *EncountersRepository) Update(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Save(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EncountersRepository) GetByIds(ids []int64) ([]model.Encounter, error) {
	var encounters []model.Encounter
	dbResult := repo.DatabaseConnection.Where("id IN ?", ids).Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return encounters, nil
}
