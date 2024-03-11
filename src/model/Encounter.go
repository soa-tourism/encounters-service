package model

import (
	"errors"
	"math"
	"strings"
)

// EncounterStatus represents the status of an encounter.
type EncounterStatus int

const (
	Draft EncounterStatus = iota
	Archived
	Published
)

// EncounterType represents the type of an encounter.
type EncounterType int

const (
	Social EncounterType = iota
	Location
	Misc
)

// Encounter represents an encounter in the explorer system.
type Encounter struct {
	AuthorId    int64
	Name        string
	Description string
	Xp          int
	Status      EncounterStatus
	Type        EncounterType
	Latitude    float64
	Longitude   float64
}

// NewEncounter creates a new Encounter with the specified parameters.
func NewEncounter(authorID int64, name, description string, xp int, encounterType EncounterType, status EncounterStatus, latitude, longitude float64) (Encounter, error) {
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
	}, nil
}

// Copy creates a copy of the given Encounter.
func (e Encounter) Copy() Encounter {
	return Encounter{
		AuthorId:    e.AuthorId,
		Name:        e.Name,
		Description: e.Description,
		Xp:          e.Xp,
		Status:      e.Status,
		Type:        e.Type,
		Latitude:    e.Latitude,
		Longitude:   e.Longitude,
	}
}

// IsValid checks if the Encounter has valid parameters.
func (e Encounter) IsValid() bool {
	return isValid(e.Name, e.Description, e.AuthorId, e.Xp, e.Longitude, e.Latitude, e.Status)
}

func isValid(name, description string, authorID int64, xp int, longitude, latitude float64, status EncounterStatus) bool {
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

func isStatusValid(status EncounterStatus) bool {
	return status != Archived
}

// GetDistanceFromEncounter calculates the distance between two points on the Earth's surface using Haversine formula.
func (e Encounter) GetDistanceFromEncounter(longitude, latitude float64) float64 {
	if latitude == e.Latitude && longitude == e.Longitude {
		return 0
	}
	return math.Acos(math.Sin(math.Pi/180*e.Latitude)*math.Sin(math.Pi/180*latitude)+
		math.Cos(math.Pi/180*e.Latitude)*math.Cos(math.Pi/180*latitude)*math.Cos(math.Pi/180*e.Longitude-math.Pi/180*longitude)) * 6371000
}

// IsCloseEnough checks if the given coordinates are close enough to the Encounter.
func (e Encounter) IsCloseEnough(longitude, latitude float64) bool {
	return e.GetDistanceFromEncounter(longitude, latitude) <= 1000
}
