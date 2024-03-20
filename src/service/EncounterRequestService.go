package service

import (
	"encounters-service/dto"
	repository "encounters-service/repositories"
	"fmt"
)

type EncounterRequestService struct {
	Repo *repository.EncounterRequestRepository
}

func (service EncounterRequestService) AcceptRequest(id int64) (dto.EncounterRequestDto, error) {
	request, err := service.Repo.AcceptRequest(id)
	if err != nil {
		return dto.EncounterRequestDto{}, fmt.Errorf("failed to accept encounter request: %v", err)
	}
	encRepo := repository.EncountersRepository{
		DatabaseConnection: service.Repo.DatabaseConnection,
	}
	enc, _ := encRepo.Get(request.EncounterId)
	enc.Status = 2
	encRepo.Update(&enc)
	request.Encounter = enc
	return dto.CreateEncounterRequestDto(request), nil
}

func (service EncounterRequestService) RejectRequest(id int64) (dto.EncounterRequestDto, error) {
	request, err := service.Repo.RejectRequest(id)
	if err != nil {
		return dto.EncounterRequestDto{}, fmt.Errorf("failed to reject encounter request: %v", err)
	}
	return dto.CreateEncounterRequestDto(request), nil
}

func (service *EncounterRequestService) Create(encounterRequestDto dto.EncounterRequestDto) (dto.EncounterRequestDto, error) {
	requestDto := encounterRequestDto.GetEncounterRequest()
	requestDto.Status = 0
	err := service.Repo.Create(&requestDto)
	if err != nil {
		return dto.EncounterRequestDto{}, fmt.Errorf("failed to create encounter request: %v", err)
	}
	return dto.CreateEncounterRequestDto(requestDto), nil
}

func (service *EncounterRequestService) GetAll() ([]dto.EncounterRequestDto, error) {
	encounterRequests, err := service.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch encounter requests: %v", err)
	}

	dtoList := make([]dto.EncounterRequestDto, 0, len(encounterRequests))

	for _, request := range encounterRequests {
		dto := dto.CreateEncounterRequestDto(request)
		dtoList = append(dtoList, dto)
	}

	return dtoList, nil
}
