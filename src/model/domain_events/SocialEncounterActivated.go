package domainevents

import (
	"encounters-service/abstractions"
	"time"
)

// SocialEncounterActivated represents a social encounter activation event in the explorer system.
type SocialEncounterActivated struct {
	abstractions.BaseDomainEvent
	ActivationDate time.Time `json:"activationDate"`
	TouristID      int64     `json:"touristId"`
}

// NewSocialEncounterActivated creates a new SocialEncounterActivated event with the specified parameters.
func NewSocialEncounterActivated(aggregateID, touristID int64, activationDate time.Time) SocialEncounterActivated {
	return SocialEncounterActivated{
		BaseDomainEvent: abstractions.NewBaseDomainEvent(aggregateID),
		TouristID:       touristID,
		ActivationDate:  activationDate,
	}
}
