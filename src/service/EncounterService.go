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
func (s EncounterService) Create(encounterDto dto.EncounterDto, checkpointId int64, isSecretPrerequisite bool, userId int64) (dto.EncounterDto, error) {
	encounter := encounterDto.GetEncounter()
	if encounter.AuthorId != userId {
		return dto.CreateEncounterDto(encounter), fmt.Errorf(fmt.Sprintf("Author is not ok"))
	}
	if encounter.IsValid() {
		s.Repo.Create(&encounter)
		return dto.CreateEncounterDto(encounter), nil
	}
	return dto.CreateEncounterDto(encounter), fmt.Errorf(fmt.Sprintf("Encounter is not valid"))
}

// TODO: proveriti checkpoint i u njemu obrisati encounter u glavnoj aplikaciji
func (s EncounterService) Delete(encounterId int64, userId int64) {
	encounter, _ := s.Get(encounterId)
	if encounter.AuthorId != userId {
		return
	}
	s.Repo.Delete(encounterId)
}

func (s EncounterService) Get(encounterId int64) (dto.EncounterDto, error) {
	encounter, err := s.Repo.Get(encounterId)
	if err != nil {
		return dto.CreateEncounterDto(encounter), err
	}
	return dto.CreateEncounterDto(encounter), nil
}

func (s EncounterService) Update(encounter dto.EncounterDto, userId int64) (dto.EncounterDto, error) {
	if userId != encounter.AuthorId {
		return encounter, fmt.Errorf(fmt.Sprintf("Not the author of the encounter"))
	}
	enc := encounter.GetEncounter()
	if !enc.IsValid() {
		return encounter, fmt.Errorf(fmt.Sprintf("Not a valid encounter"))
	}
	err := s.Repo.Update(&enc)
	if err != nil {

		return encounter, fmt.Errorf(fmt.Sprintf("Error updating encounter"))
	}
	return dto.CreateEncounterDto(enc), nil
}
