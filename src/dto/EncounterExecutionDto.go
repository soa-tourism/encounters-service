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

func CreateEncounterExecutionDto(enc model.EncounterExecution) EncounterExecutionDto {
	return EncounterExecutionDto{
		Id:           enc.Id,
		EncounterId:  enc.EncounterId,
		EncounterDto: CreateEncounterDto(enc.Encounter),
		TouristId:    enc.TouristId,
		Status:       ExecutionStatusStringConversion(enc.Status),
		StartTime:    enc.StartTime,
	}
}

func (executionDto EncounterExecutionDto) GetEncounterExecution() model.EncounterExecution {
	return model.EncounterExecution{
		Id:          executionDto.Id,
		EncounterId: executionDto.EncounterId,
		TouristId:   executionDto.TouristId,
		Encounter:   executionDto.EncounterDto.GetEncounter(),
		Status:      ExecutionStatusNumberConversion(executionDto.Status),
		StartTime:   executionDto.StartTime,
	}
}

func ExecutionStatusStringConversion(number int) string {
	if number == 0 {
		return "Pending"
	}
	if number == 1 {
		return "Completed"
	}
	if number == 3 {
		return "Abandoned"
	}
	return "Active"
}

func ExecutionStatusNumberConversion(status string) int {
	if status == "Pending" {
		return 0
	}
	if status == "Completed" {
		return 1
	}
	if status == "Abandoned" {
		return 3
	}
	return 2
}
