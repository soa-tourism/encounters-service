package service

import (
	"encounters-service/dto"
	repository "encounters-service/repositories"
	"fmt"
)

type EncounterRequestService struct {
	Repo *repository.EncounterRequestDatabaseRepository
}

func (s EncounterRequestService) AcceptRequest(id int64) (dto.EncounterRequestDto, error) {
	request, err := s.Repo.AcceptRequest(id)
	if err != nil {
		return dto.EncounterRequestDto{}, fmt.Errorf("failed to accept request: %v", err)
	}
	return dto.CreateEncounterRequestDto(request), nil
}

func (s EncounterRequestService) RejectRequest(id int64) (dto.EncounterRequestDto, error) {
	request, err := s.Repo.RejectRequest(id)
	if err != nil {
		return dto.EncounterRequestDto{}, fmt.Errorf("failed to reject request: %v", err)
	}
	return dto.CreateEncounterRequestDto(request), nil
}

func (s EncounterRequestService) Create(encounterRequestDto dto.EncounterRequestDto, encounterId int64, touristId int64) (dto.EncounterRequestDto, error) {
	requestDto := encounterRequestDto.GetEncounterRequest()

	if requestDto.EncounterId != encounterId {
		return dto.CreateEncounterRequestDto(requestDto), fmt.Errorf("Encounter not valid")
	}

	if requestDto.TouristId != touristId {
		return dto.CreateEncounterRequestDto(requestDto), fmt.Errorf("Tourist not valid")
	}
	s.Repo.Create(&requestDto)
	return dto.CreateEncounterRequestDto(requestDto), nil
}
