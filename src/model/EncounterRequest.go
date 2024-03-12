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
	Status      RequestStatus
}

func NewEncounterRequest(encounterId int64, requestStatus RequestStatus, touristId int64) *EncounterRequest {
	return &EncounterRequest{
		EncounterId: encounterId,
		Status:      requestStatus,
		TouristId:   touristId,
	}
}

func (er *EncounterRequest) AcceptRequest() {
	er.Status = Accepted
}

func (er *EncounterRequest) RejectRequest() {
	er.Status = Rejected
}
