package dto

import "encounters-service/model"

type EncounterRequestDto struct {
	Id          int64  `json:"id"`
	EncounterId int64  `json:"encounterId"`
	TouristId   int64  `json:"touristId"`
	Status      string `json:"status"`
}

func CreateEncounterRequestDto(enc model.EncounterRequest) EncounterRequestDto {
	return EncounterRequestDto{
		Id:          enc.Id,
		EncounterId: enc.EncounterId,
		TouristId:   enc.TouristId,
		Status:      RequestStatusStringConversion(enc.Status),
	}
}

func (request EncounterRequestDto) GetEncounterRequest() model.EncounterRequest {
	return model.EncounterRequest{
		Id:          request.Id,
		EncounterId: request.EncounterId,
		TouristId:   request.TouristId,
		Status:      RequestStatusNumberConversion(request.Status),
	}
}

func RequestStatusStringConversion(number int) string {
	if number == 0 {
		return "OnHold"
	}
	if number == 1 {
		return "Accepted"
	}
	return "Rejected"
}

func RequestStatusNumberConversion(status string) int {
	if status == "OnHold" {
		return 0
	}
	if status == "Accepted" {
		return 1
	}
	return 2
}
