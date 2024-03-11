package domainevents

import (
	"encounters-service/abstractions"
	"time"
)

// SocialEncounterCreated represents a social encounter creation event in the explorer system.
type SocialEncounterCreated struct {
	abstractions.BaseDomainEvent
	DateOfCreation time.Time `json:"dateOfCreation"`
	RequiredPeople int       `json:"requiredPeople"`
	Range          float64   `json:"range"`
}

// NewSocialEncounterCreated creates a new SocialEncounterCreated event with the specified parameters.
func NewSocialEncounterCreated(aggregateID int64, dateOfCreation time.Time, requiredPeople int, rangeValue float64) SocialEncounterCreated {
	return SocialEncounterCreated{
		BaseDomainEvent: abstractions.NewBaseDomainEvent(aggregateID),
		DateOfCreation:  dateOfCreation,
		RequiredPeople:  requiredPeople,
		Range:           rangeValue,
	}
}
