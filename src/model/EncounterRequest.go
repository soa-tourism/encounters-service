package model

type RequestStatus int

const (
	OnHold RequestStatus = iota
	Accepted
	Rejected
)

type EncounterRequest struct {
	Id          int64
	EncounterId int64
	TouristId   int64
	Status      int //OnHold, Accepted, Rejected
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
