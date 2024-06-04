package orchestrator

import (
	"encounters-service/proto/encounter"
	repository "encounters-service/repositories"
	events "encounters-service/saga/create_encounter"
	saga "encounters-service/saga/messaging"
	"fmt"
	"strconv"
)

type CreateEncounterOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
	database         repository.EncountersRepository
}

func NewCreateEncounterOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber, database repository.EncountersRepository) (*CreateEncounterOrchestrator, error) {
	o := &CreateEncounterOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
		database:         database,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateEncounterOrchestrator) Start(request *encounter.CreateRequest, encounterId int64) error {
	event := &events.CreateEncounterCommand{
		Type: events.UpdateCheckpoint,
		Encounter: events.EncounterCreatedEvent{
			EncounterId:          encounterId,
			CheckpointId:         request.CheckpointId,
			IsSecretPrerequisite: request.IsSecretPrerequisite,
		},
	}
	fmt.Println("CH UPDATE (encounterID = " + strconv.FormatInt(encounterId, 10))
	return o.commandPublisher.Publish(event)
}

func (o *CreateEncounterOrchestrator) handle(reply *events.CreateEncounterReply) {
	switch reply.Type {
	case events.CheckpointUpdated:
		{
			return
		}
	default:
		{
			fmt.Println("Deleting encounter with ID " + strconv.FormatInt(reply.Encounter.EncounterId, 10))
			o.database.Delete(reply.Encounter.EncounterId)
			return
		}
	}
}
