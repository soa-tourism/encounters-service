package service

import (
	"encounters-service/dto"
	repository "encounters-service/repositories"
	"fmt"
)

type EncounterService struct {
	Repo *repository.EncountersRepository
}

// TODO: dodati encounter id u checkpoint nakon upisivanja encountera u bazu
func (s EncounterService) Create(encounterDto dto.EncounterDto) (dto.EncounterDto, error) {
	encounter := encounterDto.GetEncounter()
	encounter.Status = 0
	if encounter.IsValid() {
		s.Repo.Create(&encounter)
		return dto.CreateEncounterDto(encounter), nil
	}
	return dto.CreateEncounterDto(encounter), fmt.Errorf("encounter is not valid")
}

func (s EncounterService) AddEncounterToExecution(executionDto dto.EncounterExecutionDto) dto.EncounterExecutionDto {
	executionDto.EncounterDto, _ = s.Get(executionDto.EncounterId)
	return executionDto
}

func (s EncounterService) AddEncountersToExecution(executionDtos []dto.EncounterExecutionDto) []dto.EncounterExecutionDto {
	list := make([]dto.EncounterExecutionDto, 0)
	for _, exe := range executionDtos {
		list = append(list, s.AddEncounterToExecution(exe))
	}
	return list
}

func (s EncounterService) GetAll() ([]dto.EncounterDto, error) {
	list := make([]dto.EncounterDto, 0)
	encounters, _ := s.Repo.GetPaged(0, 0)
	for _, encounter := range encounters {
		list = append(list, dto.CreateEncounterDto(encounter))
	}
	return list, nil
}

// TODO: proveriti checkpoint i u njemu obrisati encounter u glavnoj aplikaciji
func (s EncounterService) Delete(encounterId int64) {
	s.Repo.Delete(encounterId)
}

func (s EncounterService) Get(encounterId int64) (dto.EncounterDto, error) {
	encounter, err := s.Repo.Get(encounterId)
	if err != nil {
		return dto.CreateEncounterDto(encounter), err
	}
	return dto.CreateEncounterDto(encounter), nil
}

func (s EncounterService) Update(encounter dto.EncounterDto) (dto.EncounterDto, error) {
	enc := encounter.GetEncounter()
	if !enc.IsValid() {
		return encounter, fmt.Errorf("not a valid encounter")
	}
	err := s.Repo.Update(&enc)
	if err != nil {

		return encounter, fmt.Errorf("error updating encounter")
	}
	return dto.CreateEncounterDto(enc), nil
}
