package dto

import (
	"encounters-service/model"
	"time"
)

type EncounterExecutionDto struct {
	Id           int64        `json:"id"`
	EncounterId  int64        `json:"encounterId"`
	EncounterDto EncounterDto `json:"encounterDto"`
	TouristId    int64        `json:"touristId"`
	Status       string       `json:"status"`
	StartTime    time.Time    `json:"startTime"`
}

func StatusStringConversion(number int) string {
	if number == 0 {
		return "Draft"
	}
	if number == 1 {
		return "Archived"
	}
	return "Published"
}

func StatusNumberConversion(status string) int {
	if status == "Draft" {
		return 0
	}
	if status == "Archived" {
		return 1
	}
	return 2
}

func CreateEncounterExecutionDto(enc model.EncounterExecution) EncounterExecutionDto {
	return EncounterExecutionDto{
		Id:           enc.Id,
		EncounterId:  enc.EncounterId,
		EncounterDto: CreateEncounterDto(enc.Encounter),
		TouristId:    enc.TouristId,
		Status:       StatusStringConversion(int(enc.Status)),
		StartTime:    enc.StartTime,
	}
}

func (executionDto EncounterExecutionDto) GetEncounterExecution() model.EncounterExecution {
	return model.EncounterExecution{
		Id:          executionDto.Id,
		EncounterId: executionDto.EncounterId,
		TouristId:   executionDto.TouristId,
		Encounter:   executionDto.EncounterDto.GetEncounter(),
		Status:      model.EncounterExecutionStatus(StatusNumberConversion(executionDto.Status)),
		StartTime:   executionDto.StartTime,
	}
}
