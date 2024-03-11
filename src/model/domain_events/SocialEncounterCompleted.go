package domainevents

import (
	"encounters-service/abstractions"
	"time"
)

// SocialEncounterCompleted represents a social encounter completion event in the explorer system.
type SocialEncounterCompleted struct {
	abstractions.BaseDomainEvent
	CompletionDate time.Time `json:"completionDate"`
	TouristId      int64     `json:"touristId"`
}

// NewSocialEncounterCompleted creates a new SocialEncounterCompleted event with the specified parameters.
func NewSocialEncounterCompleted(aggregateID int64, completionDate time.Time, touristID int64) SocialEncounterCompleted {
	return SocialEncounterCompleted{
		BaseDomainEvent: abstractions.NewBaseDomainEvent(aggregateID),
		CompletionDate:  completionDate,
		TouristId:       touristID,
	}
}
