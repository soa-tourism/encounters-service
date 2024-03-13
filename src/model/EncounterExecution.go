package model

import (
	"encounters-service/abstractions"
	domainevents "encounters-service/model/domain_events"
	"errors"
	"math"
	"time"
)

// EncounterExecution represents an encounter execution in the explorer system.
type EncounterExecution struct {
	Id          int64                      `json:"id" gorm:"primaryKey"`
	EncounterId int64                      `json:"encounterId" gorm:"foreignKey:Id;references:Id"`
	Encounter   Encounter                  `json:"encounter"`
	TouristId   int64                      `json:"touristId"`
	Status      int                        `json:"status"` //Pending, Completed, Active, Abandoned
	StartTime   time.Time                  `json:"startTime"`
	EndTime     time.Time                  `json:"endTime"`
	Changes     []abstractions.DomainEvent `json:"-" gorm:"type:jsonb;"`
	Version     int64                      `json:"-"`
}

// NewEncounterExecution creates a new EncounterExecution with the specified parameters.
func NewEncounterExecution(encounterID int64, touristID int64, status int, startTime, endTime time.Time) (EncounterExecution, error) {
	ee := EncounterExecution{
		EncounterId: encounterID,
		TouristId:   touristID,
		Status:      status,
		StartTime:   startTime,
		EndTime:     endTime,
		Version:     0,
		Changes:     make([]abstractions.DomainEvent, 0),
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
	ee.Status = 2
	ee.StartTime = time.Now()
	if ee.Encounter.Type == 0 {
		ee.Causes(domainevents.NewSocialEncounterActivated(ee.Id, ee.TouristId, time.Now()))
	}
}

// Abandon sets the status to Abandoned.
func (ee *EncounterExecution) Abandon() {
	ee.Status = 3
}

// Complete sets the status to Completed and updates the end time.
func (ee *EncounterExecution) Complete() {
	if ee.Encounter.Type == 0 && ee.Status == 2 {
		ee.Causes(domainevents.NewSocialEncounterCompleted(ee.Id, time.Now(), ee.TouristId))
	}
	ee.Status = 2
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
