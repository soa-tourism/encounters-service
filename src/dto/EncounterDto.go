package dto

import "encounters-service/model"

type EncounterDto struct {
	Id                int64   `json:"Id"`
	AuthorId          int64   `json:"AuthorId"`
	Name              string  `json:"Name"`
	Description       string  `json:"Description"`
	Xp                int     `json:"XP"`
	Status            string  `json:"Status"`
	Type              string  `json:"Type"`
	Latitude          float64 `json:"Latitude"`
	Longitude         float64 `json:"Longitude"`
	LocationLongitude float64 `json:"LocationLongitude"`
	LocationLatitude  float64 `json:"LocationLatitude"`
	Image             string  `json:"Image"`
	Range             float64 `json:"Range"`
	RequiredPeople    int     `json:"RequiredPeople"`
	ActiveTouristsIds []int32 `json:"ActiveTouristsIds"`
}

func CreateEncounterDto(enc model.Encounter) EncounterDto {
	encounter := EncounterDto{
		Id:                enc.Id,
		AuthorId:          enc.AuthorId,
		Name:              enc.Name,
		Description:       enc.Description,
		Xp:                enc.Xp,
		Status:            EncounterStatusStringConversion(enc.Status),
		Type:              EncounterTypeStringConversion(enc.Type),
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
		Status:            EncounterStatusNumberConversion(dto.Status),
		Type:              EncounterTypeNumberConversion(dto.Type),
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

func EncounterTypeStringConversion(number int) string {
	if number == 0 {
		return "Social"
	}
	if number == 1 {
		return "Location"
	}
	return "Misc"
}

func EncounterTypeNumberConversion(status string) int {
	if status == "Social" {
		return 0
	}
	if status == "Location" {
		return 1
	}
	return 2
}

func EncounterStatusStringConversion(number int) string {
	if number == 0 {
		return "Draft"
	}
	if number == 1 {
		return "Archived"
	}
	return "Published"
}

func EncounterStatusNumberConversion(status string) int {
	if status == "Draft" {
		return 0
	}
	if status == "Archived" {
		return 1
	}
	return 2
}
