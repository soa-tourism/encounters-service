package domainevents

import (
	"encounters-service/abstractions"
	"time"
)

// SocialEncounterRangeChecked represents a social encounter range check event in the explorer system.
type SocialEncounterRangeChecked struct {
	abstractions.BaseDomainEvent
	RangeCheckedDate  time.Time `json:"rangeCheckedDate"`
	ActiveTouristsIDs []int     `json:"activeTouristsIds"`
}

// NewSocialEncounterRangeChecked creates a new SocialEncounterRangeChecked event with the specified parameters.
func NewSocialEncounterRangeChecked(aggregateID int64, activeTouristsIDs []int, rangeCheckedDate time.Time) SocialEncounterRangeChecked {
	return SocialEncounterRangeChecked{
		BaseDomainEvent:   abstractions.NewBaseDomainEvent(aggregateID),
		RangeCheckedDate:  rangeCheckedDate,
		ActiveTouristsIDs: activeTouristsIDs,
	}
}
