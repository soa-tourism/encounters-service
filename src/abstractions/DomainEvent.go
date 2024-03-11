package abstractions

// DomainEvent represents a domain event in the explorer system.
type DomainEvent interface {
	AggregateID() int64
}

// BaseDomainEvent provides a basic implementation of the DomainEvent interface.
type BaseDomainEvent struct {
	AggregateIDField int64 `json:"aggregateId"`
}

// AggregateID returns the aggregate ID of the domain event.
func (bde BaseDomainEvent) AggregateID() int64 {
	return bde.AggregateIDField
}

// NewBaseDomainEvent creates a new BaseDomainEvent with the specified aggregate ID.
func NewBaseDomainEvent(aggregateID int64) BaseDomainEvent {
	return BaseDomainEvent{AggregateIDField: aggregateID}
}
