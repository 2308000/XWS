package application

import (
	"accommodation_booking/common/domain"
	events "accommodation_booking/common/saga/create_profile"
	saga "accommodation_booking/common/saga/messaging"
)

type CreateProfileOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateProfileOrchestrator, error) {
	o := &CreateProfileOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateProfileOrchestrator) Start(profile *domain.Profile) error {
	event := &events.CreateProfileCommand{
		Type:    events.CreateProfile,
		Profile: *profile,
	}
	return o.commandPublisher.Publish(event)
}

func (o *CreateProfileOrchestrator) handle(reply *events.CreateProfileReply) {
	command := events.CreateProfileCommand{Profile: reply.Profile}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateProfileOrchestrator) nextCommandType(reply events.CreateProfileReplyType) events.CreateProfileCommandType {
	switch reply {
	case events.ProfileNotCreated:
		return events.RollbackCreatedProfile
	default:
		return events.UnknownCommand
	}
}
