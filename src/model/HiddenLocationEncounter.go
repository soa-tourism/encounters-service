package model

import (
	"math"
)

type HiddenLocationEncounter struct {
	Id                int64     `json:"id" gorm:"primaryKey;foreignKey:Id"`
	Encounter         Encounter `json:"encounter"`
	LocationLongitude float64   `json:"locationLongitude"`
	LocationLatitude  float64   `json:"locationLatitude"`
	Image             string    `json:"image"`
	Range             float64   `json:"range"`
}

func NewHiddenLocationEncounter(encounter Encounter, locationLongitude, locationLatitude float64, image string, rangeVal float64) *HiddenLocationEncounter {
	return &HiddenLocationEncounter{
		Encounter:         encounter,
		LocationLongitude: locationLongitude,
		LocationLatitude:  locationLatitude,
		Image:             image,
		Range:             rangeVal,
	}
}

func (hle *HiddenLocationEncounter) CheckIfInRangeLocation(touristLongitude, touristLatitude float64) bool {
	distance := math.Acos(math.Sin(math.Pi/180*hle.LocationLatitude)*math.Sin(math.Pi/180*touristLatitude)+math.Cos(math.Pi/180*hle.LocationLatitude)*math.Cos(math.Pi/180*touristLatitude)*math.Cos(math.Pi/180*hle.LocationLongitude-math.Pi/180*touristLongitude)) * 6371000
	if distance <= hle.Range {
		return true
	}
	return false
}

func (hle *HiddenLocationEncounter) CheckIfLocationFound(touristLongitude, touristLatitude float64) bool {
	distance := math.Acos(math.Sin(math.Pi/180*hle.LocationLatitude)*math.Sin(math.Pi/180*touristLatitude)+math.Cos(math.Pi/180*hle.LocationLatitude)*math.Cos(math.Pi/180*touristLatitude)*math.Cos(math.Pi/180*hle.LocationLongitude-math.Pi/180*touristLongitude)) * 6371000
	if distance <= 5.0 {
		return true
	}
	return false
}
