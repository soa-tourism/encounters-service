package abstractions

type DomainEvent interface {
	AggregateID() int64
}

type BaseDomainEvent struct {
	AggregateIDField int64 `json:"aggregateId"`
}

func (bde BaseDomainEvent) AggregateID() int64 {
	return bde.AggregateIDField
}

func NewBaseDomainEvent(aggregateID int64) BaseDomainEvent {
	return BaseDomainEvent{AggregateIDField: aggregateID}
}
