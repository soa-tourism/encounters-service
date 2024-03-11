package model

import (
	"encounters-service/abstractions"
	domainevents "encounters-service/model/domain_events"
	"errors"
	"math"
	"time"
)

// EncounterExecutionStatus represents the status of an encounter execution.
type EncounterExecutionStatus int

const (
	Pending EncounterExecutionStatus = iota
	Completed
	Active
	Abandoned
)

// EncounterExecution represents an encounter execution in the explorer system.
type EncounterExecution struct {
	Id          int64
	EncounterId int64
	Encounter   Encounter
	TouristId   int64
	Status      EncounterExecutionStatus
	StartTime   time.Time
	EndTime     time.Time
	Changes     []abstractions.DomainEvent
	Version     int64
}

// NewEncounterExecution creates a new EncounterExecution with the specified parameters.
func NewEncounterExecution(encounterID int64, encounter Encounter, touristID int64, status EncounterExecutionStatus, startTime, endTime time.Time) (EncounterExecution, error) {
	ee := EncounterExecution{
		EncounterId: encounterID,
		Encounter:   encounter,
		TouristId:   touristID,
		Status:      status,
		StartTime:   startTime,
		EndTime:     endTime,
	}
	if err := ee.Validate(); err != nil {
		return EncounterExecution{}, err
	}
	return ee, nil
}

// Validate checks if the EncounterExecution has valid parameters.
func (ee *EncounterExecution) Validate() error {
	if ee.EncounterId == 0 {
		return errors.New("invalid encounter ID")
	}
	if ee.TouristId == 0 {
		return errors.New("invalid tourist")
	}
	if ee.Status == 0 {
		return errors.New("invalid execution status")
	}
	if ee.StartTime.After(time.Now()) {
		return errors.New("invalid start time")
	}
	if ee.EndTime.After(time.Now()) {
		return errors.New("invalid end time")
	}
	return nil
}

// Activate sets the status to Active and updates the start time.
func (ee *EncounterExecution) Activate() {
	ee.Status = Active
	ee.StartTime = time.Now()
	if ee.Encounter.Type == Social {
		ee.Causes(domainevents.NewSocialEncounterActivated(ee.Id, ee.TouristId, time.Now()))
	}
}

// Abandon sets the status to Abandoned.
func (ee *EncounterExecution) Abandon() {
	ee.Status = Abandoned
}

// Complete sets the status to Completed and updates the end time.
func (ee *EncounterExecution) Complete() {
	if ee.Encounter.Type == Social && ee.Status == Active {
		ee.Causes(domainevents.NewSocialEncounterCompleted(ee.Id, time.Now(), ee.TouristId))
	}
	ee.Status = Completed
	ee.EndTime = time.Now()
}

// CheckRangeDistance calculates the distance between two points on the Earth's surface using Haversine formula.
func (ee *EncounterExecution) CheckRangeDistance(touristLongitude, touristLatitude float64) float64 {
	if touristLatitude == ee.Encounter.Latitude && touristLongitude == ee.Encounter.Longitude {
		return 0
	}
	distance := math.Acos(math.Sin(math.Pi/180*ee.Encounter.Latitude)*math.Sin(math.Pi/180*touristLatitude)+
		math.Cos(math.Pi/180*ee.Encounter.Latitude)*math.Cos(math.Pi/180*touristLatitude)*math.Cos(math.Pi/180*ee.Encounter.Longitude-math.Pi/180*touristLongitude)) * 6371000
	return distance
}

// Causes adds a domain event to the Changes list and applies the event.
func (ee *EncounterExecution) Causes(event abstractions.DomainEvent) {
	ee.Changes = append(ee.Changes, event)
	ee.Apply(event)
}

// Apply increments the version when a domain event is applied.
func (ee *EncounterExecution) Apply(event abstractions.DomainEvent) {
	ee.Version++
}
