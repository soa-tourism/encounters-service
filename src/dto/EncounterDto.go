package dto

import "encounters-service/model"

type EncounterDto struct {
	Id                int64                 `json:"id"`
	AuthorId          int64                 `json:"authorId"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Xp                int                   `json:"xp"`
	Status            model.EncounterStatus `json:"status"`
	Type              model.EncounterType   `json:"type"`
	Latitude          float64               `json:"latitude"`
	Longitude         float64               `json:"longitude"`
	LocationLongitude float64               `json:"locationLongitude"`
	LocationLatitude  float64               `json:"locationLatitude"`
	Image             string                `json:"image"`
	Range             float64               `json:"range"`
	RequiredPeople    int                   `json:"requiredPeople"`
	ActiveTouristsIds []int                 `json:"activeTouristsIds"`
}

func CreateEncounterDto(enc model.Encounter) EncounterDto {
	encounter := EncounterDto{
		Id:          enc.Id,
		AuthorId:    enc.AuthorId,
		Name:        enc.Name,
		Description: enc.Description,
		Xp:          enc.Xp,
		Status:      enc.Status,
		Type:        enc.Type,
		Latitude:    enc.Latitude,
		Longitude:   enc.Longitude,
	}
	return encounter
}

func CreateSocialEncounterDto(enc model.SocialEncounter) EncounterDto {
	encounter := EncounterDto{
		Id:                enc.Id,
		AuthorId:          enc.Encounter.AuthorId,
		Name:              enc.Encounter.Name,
		Description:       enc.Encounter.Description,
		Xp:                enc.Encounter.Xp,
		Status:            enc.Encounter.Status,
		Type:              enc.Encounter.Type,
		Latitude:          enc.Encounter.Latitude,
		Longitude:         enc.Encounter.Longitude,
		RequiredPeople:    enc.RequiredPeople,
		Range:             enc.Range,
		ActiveTouristsIds: enc.ActiveTouristsIds,
	}
	return encounter
}

func CreateHiddenLocationEncounterDto(enc model.HiddenLocationEncounter) EncounterDto {
	encounter := EncounterDto{
		Id:                enc.Id,
		AuthorId:          enc.Encounter.AuthorId,
		Name:              enc.Encounter.Name,
		Description:       enc.Encounter.Description,
		Xp:                enc.Encounter.Xp,
		Status:            enc.Encounter.Status,
		Type:              enc.Encounter.Type,
		Latitude:          enc.Encounter.Latitude,
		Longitude:         enc.Encounter.Longitude,
		LocationLongitude: enc.LocationLongitude,
		LocationLatitude:  enc.LocationLatitude,
		Image:             enc.Image,
	}
	return encounter
}

func (dto EncounterDto) GetEncounter() model.Encounter {
	encounter := model.Encounter{
		Id:          dto.Id,
		AuthorId:    dto.AuthorId,
		Name:        dto.Name,
		Description: dto.Description,
		Xp:          dto.Xp,
		Status:      dto.Status,
		Type:        dto.Type,
		Longitude:   dto.Longitude,
		Latitude:    dto.Latitude,
	}
	return encounter
}
func (dto EncounterDto) GetSocialEncounter() model.SocialEncounter {
	socialEncounter := model.SocialEncounter{
		Encounter: model.Encounter{
			Id:          dto.Id,
			AuthorId:    dto.AuthorId,
			Name:        dto.Name,
			Description: dto.Description,
			Xp:          dto.Xp,
			Status:      dto.Status,
			Type:        dto.Type,
			Longitude:   dto.Longitude,
			Latitude:    dto.Latitude,
		},
		RequiredPeople:    dto.RequiredPeople,
		Range:             dto.Range,
		ActiveTouristsIds: dto.ActiveTouristsIds,
	}
	return socialEncounter
}
func (dto EncounterDto) GetHiddenLocationEncounter() model.HiddenLocationEncounter {
	hiddenLocationEncounter := model.HiddenLocationEncounter{
		Encounter: model.Encounter{
			Id:          dto.Id,
			AuthorId:    dto.AuthorId,
			Name:        dto.Name,
			Description: dto.Description,
			Xp:          dto.Xp,
			Status:      dto.Status,
			Type:        dto.Type,
			Longitude:   dto.Longitude,
			Latitude:    dto.Latitude,
		},
		LocationLongitude: dto.LocationLongitude,
		LocationLatitude:  dto.LocationLatitude,
		Image:             dto.Image,
	}
	return hiddenLocationEncounter
}
