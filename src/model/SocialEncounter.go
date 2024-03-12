package model

import (
	"database/sql/driver"
	"encoding/json"
	domainevents "encounters-service/model/domain_events"
	"errors"
	"time"
)

type SocialEncounter struct {
	Id                int64     `json:"id" gorm:"primaryKey;foreignKey:Id"`
	Encounter         Encounter `json:"encounter"`
	RequiredPeople    int       `json:"requiredPeople"`
	Range             float64   `json:"range"`
	ActiveTouristsIds []int     `json:"activeTouristsIds"`
}

func NewSocialEncounter(encounter Encounter, requiredPeople int, rangeVal float64) SocialEncounter {
	return SocialEncounter{
		Encounter:         encounter,
		RequiredPeople:    requiredPeople,
		Range:             rangeVal,
		ActiveTouristsIds: make([]int, 0),
	}
}

func (s *SocialEncounter) CreateSocialEncounter() {
	if s.Encounter.Type == Social {
		s.Encounter.Causes(domainevents.NewSocialEncounterCreated(s.Id, time.Now(), s.RequiredPeople, s.Range))
	}
}

func (s *SocialEncounter) CheckIfInRange(touristLongitude, touristLatitude float64, touristId int) int {
	s.Encounter.Causes(domainevents.NewSocialEncounterLocationUpdated(s.Id, touristId, time.Now(), s.Encounter.Longitude, s.Encounter.Latitude))
	distance := s.Encounter.GetDistanceFromEncounter(touristLongitude, touristLatitude)
	if distance > s.Range {
		s.RemoveTourist(touristId)
		return 0
	}
	s.AddTourist(touristId)
	return len(s.ActiveTouristsIds)
}

func (s *SocialEncounter) AddTourist(touristId int) {
	if !s.isTouristInActiveList(touristId) {
		s.ActiveTouristsIds = append(s.ActiveTouristsIds, touristId)
		s.Encounter.Causes(domainevents.NewSocialEncounterRangeChecked(s.Id, s.ActiveTouristsIds, time.Now()))
	}
}

func (s *SocialEncounter) RemoveTourist(touristId int) {
	if s.isTouristInActiveList(touristId) {
		for i, id := range s.ActiveTouristsIds {
			if id == touristId {
				s.ActiveTouristsIds = append(s.ActiveTouristsIds[:i], s.ActiveTouristsIds[i+1:]...)
				s.Encounter.Causes(domainevents.NewSocialEncounterRangeChecked(s.Id, s.ActiveTouristsIds, time.Now()))
				break
			}
		}
	}
}

func (s *SocialEncounter) IsRequiredPeopleNumber() bool {
	numberOfTourists := len(s.ActiveTouristsIds)
	if numberOfTourists >= s.RequiredPeople {
		s.ClearActiveTourists()
	}
	return numberOfTourists >= s.RequiredPeople
}

func (s *SocialEncounter) ClearActiveTourists() {
	s.ActiveTouristsIds = []int{}
}

func (s *SocialEncounter) isTouristInActiveList(touristId int) bool {
	for _, id := range s.ActiveTouristsIds {
		if id == touristId {
			return true
		}
	}
	return false
}

func (s *SocialEncounter) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("unsupported type for Scan")
	}

	return json.Unmarshal(bytes, s)
}

func (s SocialEncounter) Value() (driver.Value, error) {
	return json.Marshal(s)
}
