package dto

import "encounters-service/model"

type EncounterDto struct {
	Id                int64   `json:"id"`
	AuthorId          int64   `json:"authorId"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Xp                int     `json:"xp"`
	Status            int     `json:"status"`
	Type              int     `json:"type"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	LocationLongitude float64 `json:"locationLongitude"`
	LocationLatitude  float64 `json:"locationLatitude"`
	Image             string  `json:"image"`
	Range             float64 `json:"range"`
	RequiredPeople    int     `json:"requiredPeople"`
	ActiveTouristsIds []int   `json:"activeTouristsIds"`
}

func CreateEncounterDto(enc model.Encounter) EncounterDto {
	encounter := EncounterDto{
		Id:                enc.Id,
		AuthorId:          enc.AuthorId,
		Name:              enc.Name,
		Description:       enc.Description,
		Xp:                enc.Xp,
		Status:            enc.Status,
		Type:              enc.Type,
		Latitude:          enc.Latitude,
		Longitude:         enc.Longitude,
		LocationLongitude: enc.LocationLongitude,
		LocationLatitude:  enc.Latitude,
		Image:             enc.Image,
		Range:             enc.Range,
		ActiveTouristsIds: enc.ActiveTouristsIds,
		RequiredPeople:    enc.RequiredPeople,
	}
	return encounter
}

func (dto EncounterDto) GetEncounter() model.Encounter {
	encounter := model.Encounter{
		Id:                dto.Id,
		AuthorId:          dto.AuthorId,
		Name:              dto.Name,
		Description:       dto.Description,
		Xp:                dto.Xp,
		Status:            dto.Status,
		Type:              dto.Type,
		Longitude:         dto.Longitude,
		Latitude:          dto.Latitude,
		LocationLongitude: dto.LocationLongitude,
		LocationLatitude:  dto.Latitude,
		Image:             dto.Image,
		Range:             dto.Range,
		ActiveTouristsIds: dto.ActiveTouristsIds,
		RequiredPeople:    dto.RequiredPeople,
	}
	return encounter
}
