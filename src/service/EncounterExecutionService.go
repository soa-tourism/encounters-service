package service

import (
	"encounters-service/dto"
	"encounters-service/model"
	repository "encounters-service/repositories"
	"fmt"
)

type EncounterExecutionService struct {
	Repo *repository.EncountersExecutionRepository
}

func (s EncounterExecutionService) Create(encounterDto dto.EncounterExecutionDto, userId int64) (dto.EncounterExecutionDto, error) {
	fmt.Println(encounterDto)
	encounter := encounterDto.GetEncounterExecution()
	if encounter.TouristId != userId {
		return dto.CreateEncounterExecutionDto(encounter), fmt.Errorf("tourist is not ok")
	}
	err := encounter.Validate()
	if err != nil {
		return dto.CreateEncounterExecutionDto(encounter), fmt.Errorf("encounter is not valid")
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

func (s EncounterExecutionService) Update(execution dto.EncounterExecutionDto) (dto.EncounterExecutionDto, error) {
	enc := execution.GetEncounterExecution()
	err := enc.Validate()
	if err != nil {
		return dto.CreateEncounterExecutionDto(enc), fmt.Errorf("encounter is not valid")
	}
	er := s.Repo.Update(&enc)
	if er != nil {
		return execution, fmt.Errorf("error updating encounter")
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

func (s EncounterExecutionService) isTouristInRange(execution model.EncounterExecution, touristLatitude float64, touristLongitude float64) bool {
	if execution.Encounter.Type == 2 && execution.CheckRangeDistance(touristLongitude, touristLatitude) <= 300.0 {
		return true
	}
	if execution.Encounter.Type == 0 && execution.CheckRangeDistance(touristLongitude, touristLatitude) <= execution.Encounter.Range {
		return true
	}
	if execution.Encounter.Type == 1 && execution.Encounter.CheckIfInRangeLocation(touristLongitude, touristLatitude) {
		return true
	}
	return false
}

func (s EncounterExecutionService) Activate(touristId int64, touristLatitude float64, touristLongitude float64, encounterId int64) (dto.EncounterExecutionDto, error) {
	execution, _ := s.Repo.GetByEncounterAndTourist(touristId, encounterId)
	if execution.Status == 1 {
		return dto.EncounterExecutionDto{}, fmt.Errorf("encounter already completed")
	}
	if s.isTouristInRange(execution, touristLatitude, touristLongitude) {
		execution.Activate()
		s.Repo.Update(&execution)
		return dto.CreateEncounterExecutionDto(execution), nil
	}
	return dto.EncounterExecutionDto{}, fmt.Errorf("error activating encounter")
}

func (s EncounterExecutionService) GetVisibleByTour(encounterIds []int64, touristLongitude float64, touristLatitude float64, touristId int64) dto.EncounterExecutionDto {
	repo := repository.EncountersRepository{
		DatabaseConnection: s.Repo.DatabaseConnection,
	}
	encounters, _ := repo.GetByIds(encounterIds)
	if len(encounters) > 0 {
		closestEncounter := encounters[0]
		bestDistance := closestEncounter.GetDistanceFromEncounter(touristLongitude, touristLatitude)

		for _, encounter := range encounters {
			distance := encounter.GetDistanceFromEncounter(touristLongitude, touristLatitude)
			if distance < bestDistance && encounter.IsCloseEnough(touristLongitude, touristLatitude) {
				bestDistance = distance
				closestEncounter = encounter
			}
		}
		execution, err := s.Repo.GetByEncounterAndTourist(touristId, closestEncounter.Id)
		var encounterDto dto.EncounterExecutionDto
		if err != nil {
			encounterDto = s.createNewEcounterExecution(touristId, closestEncounter)
		} else {
			encounterDto = dto.CreateEncounterExecutionDto(execution)
		}
		return encounterDto
	}
	return dto.EncounterExecutionDto{}
}

func (s EncounterExecutionService) createNewEcounterExecution(touristId int64, closestEncounter model.Encounter) dto.EncounterExecutionDto {
	execution := model.EncounterExecution{
		EncounterId: closestEncounter.Id,
		Encounter:   closestEncounter,
		TouristId:   touristId,
	}
	s.Repo.Create(&execution)
	return dto.CreateEncounterExecutionDto(execution)
}

func (s EncounterExecutionService) CheckIfInRange(id int64, touristLongitude float64, touristLatitude float64, touristId int64) (dto.EncounterExecutionDto, error) {
	oldExecution, _ := s.Repo.Get(id)
	if oldExecution.Status != 2 || oldExecution.Encounter.Type != 0 {
		return dto.EncounterExecutionDto{}, fmt.Errorf("encounter is not active")
	}
	socialEncounter := oldExecution.Encounter
	socialEncounter.CheckIfInRange(touristLongitude, touristLatitude, int(touristId))
	encounterRepository := repository.EncountersRepository{
		DatabaseConnection: s.Repo.DatabaseConnection,
	}
	encounterRepository.Update(&socialEncounter)
	if socialEncounter.IsRequiredPeopleNumber() {
		allActiveSocial, _ := s.Repo.GetBySocialEncounter(socialEncounter.Id)
		for _, activeSocial := range allActiveSocial {
			if activeSocial.Status == 2 && activeSocial.Id != id {
				s.CompleteExecution(activeSocial.Id, activeSocial.TouristId, touristLatitude, touristLongitude)
			}
		}
		execution, _ := s.CompleteExecution(id, touristId, touristLatitude, touristLongitude)
		return execution, nil
	}
	return dto.CreateEncounterExecutionDto(oldExecution), nil
}

// TODO update XP level of the tourist
func (s EncounterExecutionService) CompleteExecution(id int64, touristId int64, touristLatitude float64, touristLongitude float64) (dto.EncounterExecutionDto, error) {
	var encounterExecution model.EncounterExecution
	encounterExecution, _ = s.Repo.Get(id)
	if encounterExecution.TouristId != touristId {
		return dto.EncounterExecutionDto{}, fmt.Errorf("encounter is not active")
	}
	if encounterExecution.Status != 2 {
		return dto.EncounterExecutionDto{}, fmt.Errorf("encounter is not active")
	}
	if s.isTouristInRange(encounterExecution, touristLatitude, touristLongitude) {
		encounterExecution.Complete()
		s.updateAllSocialCompleted(encounterExecution.EncounterId, encounterExecution.Encounter.Type)
		s.Repo.Update(&encounterExecution)
		return dto.CreateEncounterExecutionDto(encounterExecution), nil
	}
	return dto.EncounterExecutionDto{}, fmt.Errorf("encounter is not active")
}

func (s EncounterExecutionService) updateAllSocialCompleted(encounterId int64, t int) {
	completed := make([]model.EncounterExecution, 0)
	var list []model.EncounterExecution
	if t == 0 {
		list, _ = s.Repo.GetBySocialEncounter(encounterId)
	} else {
		return
	}
	for _, e := range list {
		completed = append(completed, e)
	}
	s.Repo.UpdateRange(completed)
}

func (s EncounterExecutionService) GetByEncounter(encounterId int64) []dto.EncounterExecutionDto {
	result, _ := s.Repo.GetByEncounter(encounterId)
	list := make([]dto.EncounterExecutionDto, 0)
	for _, e := range result {
		list = append(list, dto.CreateEncounterExecutionDto(e))
	}
	return list
}

func (s EncounterExecutionService) GetWithUpdatedLocation(encounterIds []int64, id int64, touristLongitude float64, touristLatitude float64, touristId int64) dto.EncounterExecutionDto {
	s.CheckIfInRange(id, touristLongitude, touristLatitude, touristId)
	return s.GetVisibleByTour(encounterIds, touristLongitude, touristLatitude, touristId)
}

func (s EncounterExecutionService) GetHiddenLocationEncounterWithUpdatedLocation(encounterIds []int64, id int64, touristLongitude float64, touristLatitude float64, touristId int64) dto.EncounterExecutionDto {
	s.CheckIfInRangeLocation(id, touristLongitude, touristLatitude, touristId)
	return s.GetVisibleByTour(encounterIds, touristLongitude, touristLatitude, touristId)
}

func (s EncounterExecutionService) GetActiveByTour(touristId int64, encounterIds []int64) []dto.EncounterExecutionDto {
	result, _ := s.Repo.GetActiveByTourist(touristId)
	list := make([]dto.EncounterExecutionDto, 0)
	for _, e := range result {
		for _, r := range encounterIds {
			if r == e.EncounterId {
				list = append(list, dto.CreateEncounterExecutionDto(e))
				break
			}
		}
	}
	return list
}

func (s EncounterExecutionService) CheckIfInRangeLocation(id int64, touristLongitude float64, touristLatitude float64, touristId int64) (dto.EncounterExecutionDto, error) {
	oldExecution, _ := s.Repo.Get(id)
	if oldExecution.Status != 2 || oldExecution.Encounter.Type != 1 {
		return dto.EncounterExecutionDto{}, fmt.Errorf("encounter is not active")
	}
	socialEncounter := oldExecution.Encounter
	if socialEncounter.CheckIfInRangeLocation(touristLongitude, touristLatitude) {
		encounterRepository := repository.EncountersRepository{
			DatabaseConnection: s.Repo.DatabaseConnection,
		}
		encounterRepository.Update(&socialEncounter)
		if socialEncounter.CheckIfLocationFound(touristLongitude, touristLatitude) {
			execution, _ := s.CompleteExecution(id, touristId, touristLatitude, touristLongitude)
			return execution, nil
		}
	}

	return dto.CreateEncounterExecutionDto(oldExecution), nil
}
