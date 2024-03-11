package domainevents

import (
	"encounters-service/abstractions"
	"time"
)

// SocialEncounterLocationUpdated represents a social encounter location update event in the explorer system.
type SocialEncounterLocationUpdated struct {
	abstractions.BaseDomainEvent
	DateOfUpdate time.Time `json:"dateOfUpdate"`
	TouristID    int       `json:"touristId"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
}

// NewSocialEncounterLocationUpdated creates a new SocialEncounterLocationUpdated event with the specified parameters.
func NewSocialEncounterLocationUpdated(aggregateID int64, touristID int, dateOfUpdate time.Time, longitude, latitude float64) SocialEncounterLocationUpdated {
	return SocialEncounterLocationUpdated{
		BaseDomainEvent: abstractions.NewBaseDomainEvent(aggregateID),
		TouristID:       touristID,
		DateOfUpdate:    dateOfUpdate,
		Longitude:       longitude,
		Latitude:        latitude,
	}
}
