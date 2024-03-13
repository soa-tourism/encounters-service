package model

import (
	"encounters-service/abstractions"
	domainevents "encounters-service/model/domain_events"
	"errors"
	"math"
	"strings"
	"time"
)

type Encounter struct {
	Id                int64                      `json:"id" gorm:"primaryKey"`
	AuthorId          int64                      `json:"authorId"`
	Name              string                     `json:"name"`
	Description       string                     `json:"description"`
	Xp                int                        `json:"xp"`
	Status            int                        `json:"status"` // Draft, Archived, Published
	Type              int                        `json:"type"`   // Social, Location, Misc
	Latitude          float64                    `json:"latitude"`
	Longitude         float64                    `json:"longitude"`
	RequiredPeople    int                        `json:"requiredPeople"`
	Range             float64                    `json:"range"`
	ActiveTouristsIds []int                      `json:"activeTouristsIds" gorm:"type:jsonb"`
	LocationLongitude float64                    `json:"locationLongitude"`
	LocationLatitude  float64                    `json:"locationLatitude"`
	Image             string                     `json:"image"`
	Changes           []abstractions.DomainEvent `json:"-" gorm:"type:jsonb"`
	Version           int64                      `json:"-"`
}

func NewEncounter(authorID int64, name, description string, xp int, encounterType int, status int, latitude, longitude float64) (Encounter, error) {
	if valid := isValid(name, description, authorID, xp, longitude, latitude, status); !valid {
		return Encounter{}, errors.New("invalid parameters for Encounter")
	}

	return Encounter{
		AuthorId:    authorID,
		Name:        name,
		Description: description,
		Xp:          xp,
		Status:      status,
		Type:        encounterType,
		Latitude:    latitude,
		Longitude:   longitude,
		Version:     0,
		Changes:     make([]abstractions.DomainEvent, 0),
	}, nil
}

func (e Encounter) IsValid() bool {
	return isValid(e.Name, e.Description, e.AuthorId, e.Xp, e.Longitude, e.Latitude, e.Status)
}

func isValid(name, description string, authorID int64, xp int, longitude, latitude float64, status int) bool {
	return isNameValid(name) && isDescriptionValid(description) && isXPValid(xp) &&
		isAuthorIDValid(authorID) && isLongitudeValid(longitude) && isLatitudeValid(latitude) && isStatusValid(status)
}

func isNameValid(name string) bool {
	return !isEmptyOrWhitespace(name)
}

func isDescriptionValid(description string) bool {
	return !isEmptyOrWhitespace(description)
}

func isEmptyOrWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isAuthorIDValid(authorID int64) bool {
	return authorID != 0
}

func isXPValid(xp int) bool {
	return xp >= 0
}

func isLongitudeValid(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}

func isLatitudeValid(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}

func isStatusValid(status int) bool {
	return status != 1
}

func (e Encounter) GetDistanceFromEncounter(longitude, latitude float64) float64 {
	if latitude == e.Latitude && longitude == e.Longitude {
		return 0
	}
	return math.Acos(math.Sin(math.Pi/180*e.Latitude)*math.Sin(math.Pi/180*latitude)+
		math.Cos(math.Pi/180*e.Latitude)*math.Cos(math.Pi/180*latitude)*math.Cos(math.Pi/180*e.Longitude-math.Pi/180*longitude)) * 6371000
}

func (e Encounter) IsCloseEnough(longitude, latitude float64) bool {
	return e.GetDistanceFromEncounter(longitude, latitude) <= 1000
}

func (e Encounter) MakeEncounterPublished() {
	e.Status = 2
}

func (hle *Encounter) CheckIfInRangeLocation(touristLongitude, touristLatitude float64) bool {
	if hle.Status != 1 {
		return false
	}
	distance := math.Acos(math.Sin(math.Pi/180*hle.LocationLatitude)*math.Sin(math.Pi/180*touristLatitude)+math.Cos(math.Pi/180*hle.LocationLatitude)*math.Cos(math.Pi/180*touristLatitude)*math.Cos(math.Pi/180*hle.LocationLongitude-math.Pi/180*touristLongitude)) * 6371000
	if distance <= hle.Range {
		return true
	}
	return false
}

func (hle *Encounter) CheckIfLocationFound(touristLongitude, touristLatitude float64) bool {
	if hle.Status != 1 {
		return false
	}
	distance := math.Acos(math.Sin(math.Pi/180*hle.LocationLatitude)*math.Sin(math.Pi/180*touristLatitude)+math.Cos(math.Pi/180*hle.LocationLatitude)*math.Cos(math.Pi/180*touristLatitude)*math.Cos(math.Pi/180*hle.LocationLongitude-math.Pi/180*touristLongitude)) * 6371000
	if distance <= 5.0 {
		return true
	}
	return false
}

func (s *Encounter) CheckIfInRange(touristLongitude, touristLatitude float64, touristId int) int {
	if s.Status != 0 {
		return -1
	}
	s.Causes(domainevents.NewSocialEncounterLocationUpdated(s.Id, touristId, time.Now(), s.Longitude, s.Latitude))
	distance := s.GetDistanceFromEncounter(touristLongitude, touristLatitude)
	if distance > s.Range {
		s.RemoveTourist(touristId)
		return 0
	}
	s.AddTourist(touristId)
	return len(s.ActiveTouristsIds)
}

func (s *Encounter) AddTourist(touristId int) {
	if s.Status != 0 {
		return
	}
	if !s.isTouristInActiveList(touristId) {
		s.ActiveTouristsIds = append(s.ActiveTouristsIds, touristId)
		s.Causes(domainevents.NewSocialEncounterRangeChecked(s.Id, s.ActiveTouristsIds, time.Now()))
	}
}

func (s *Encounter) RemoveTourist(touristId int) {
	if s.Status != 0 {
		return
	}
	if s.isTouristInActiveList(touristId) {
		for i, id := range s.ActiveTouristsIds {
			if id == touristId {
				s.ActiveTouristsIds = append(s.ActiveTouristsIds[:i], s.ActiveTouristsIds[i+1:]...)
				s.Causes(domainevents.NewSocialEncounterRangeChecked(s.Id, s.ActiveTouristsIds, time.Now()))
				break
			}
		}
	}
}

func (s *Encounter) IsRequiredPeopleNumber() bool {
	if s.Status != 0 {
		return false
	}
	numberOfTourists := len(s.ActiveTouristsIds)
	if numberOfTourists >= s.RequiredPeople {
		s.ClearActiveTourists()
	}
	return numberOfTourists >= s.RequiredPeople
}

func (s *Encounter) ClearActiveTourists() {
	if s.Status == 0 {
		s.ActiveTouristsIds = []int{}
	}
}

func (s *Encounter) isTouristInActiveList(touristId int) bool {
	if s.Status != 0 {
		return false
	}
	for _, id := range s.ActiveTouristsIds {
		if id == touristId {
			return true
		}
	}
	return false
}

func (ee *Encounter) Causes(event abstractions.DomainEvent) {
	ee.Changes = append(ee.Changes, event)
	ee.Apply(event)
}

func (ee *Encounter) Apply(event abstractions.DomainEvent) {
	ee.Version++
}
