package model

type EncounterRequest struct {
	Id          int64
	EncounterId int64 `gorm:"foreignKey:Id;references:Id"`
	TouristId   int64
	Status      int // 0 for OnHold, 1 for Accepted, 2 for Rejected
	Encounter   Encounter
}

func NewEncounterRequest(encounterId int64, requestStatus int, touristId int64) *EncounterRequest {
	return &EncounterRequest{
		EncounterId: encounterId,
		Status:      requestStatus,
		TouristId:   touristId,
	}
}

func (er *EncounterRequest) AcceptRequest() {
	er.Status = 1
}

func (er *EncounterRequest) RejectRequest() {
	er.Status = 2
}
