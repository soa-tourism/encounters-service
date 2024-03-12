package service

import (
	"encounters-service/dto"
	repository "encounters-service/repositories"
	"fmt"
)

type EncounterExecutionService struct {
	Repo *repository.EncountersExecutionRepository
}

func (s EncounterExecutionService) Create(encounterDto dto.EncounterExecutionDto, userId int64) (dto.EncounterExecutionDto, error) {
	encounter := encounterDto.GetEncounterExecution()
	if encounter.TouristId != userId {
		return dto.CreateEncounterExecutionDto(encounter), fmt.Errorf(fmt.Sprintf("Tourist is not ok"))
	}
	err := encounter.Validate()
	if err != nil {
		return dto.CreateEncounterExecutionDto(encounter), fmt.Errorf(fmt.Sprintf("Encounter is not valid"))
	}
	s.Repo.Create(&encounter)
	return dto.CreateEncounterExecutionDto(encounter), nil
}

func (s EncounterExecutionService) Delete(encounterId int64, touristId int64) {
	execution, _ := s.Get(encounterId)
	if execution.TouristId != touristId {
		return
	}
	s.Repo.Delete(encounterId)
}

func (s EncounterExecutionService) Get(executionId int64) (dto.EncounterExecutionDto, error) {
	execution, err := s.Repo.Get(executionId)
	if err != nil {
		return dto.CreateEncounterExecutionDto(execution), err
	}
	return dto.CreateEncounterExecutionDto(execution), nil
}

func (s EncounterExecutionService) Update(execution dto.EncounterExecutionDto, userId int64) (dto.EncounterExecutionDto, error) {
	if userId != execution.TouristId {
		return execution, fmt.Errorf(fmt.Sprintf("Not the tourist of the execution"))
	}
	enc := execution.GetEncounterExecution()
	err := enc.Validate()
	if err != nil {
		return dto.CreateEncounterExecutionDto(enc), fmt.Errorf(fmt.Sprintf("Encounter is not valid"))
	}
	er := s.Repo.Update(&enc)
	if er != nil {
		return execution, fmt.Errorf(fmt.Sprintf("Error updating encounter"))
	}
	return dto.CreateEncounterExecutionDto(enc), nil
}

func (s EncounterExecutionService) GetAllByTourist(touristId int64) []dto.EncounterExecutionDto {
	executions, _ := s.Repo.GetAllByTourist(touristId)
	executionDtos := make([]dto.EncounterExecutionDto, 0)
	for _, exec := range executions {
		executionDtos = append(executionDtos, dto.CreateEncounterExecutionDto(exec))
	}
	return executionDtos
}

func (s EncounterExecutionService) GetAllCompletedByTourist(touristId int64) []dto.EncounterExecutionDto {
	executions, _ := s.Repo.GetAllCompletedByTourist(touristId)
	executionDtos := make([]dto.EncounterExecutionDto, 0)
	for _, exec := range executions {
		executionDtos = append(executionDtos, dto.CreateEncounterExecutionDto(exec))
	}
	return executionDtos
}
