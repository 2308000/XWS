package application

import (
	"accommodation_booking/common/domain"
	saga "accommodation_booking/common/saga/messaging"
	events "accommodation_booking/common/saga/update_profile"
)

type UpdateProfileOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUpdateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileOrchestrator, error) {
	o := &UpdateProfileOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UpdateProfileOrchestrator) Start(newProfile *domain.Profile, oldProfile *domain.Profile) error {
	event := &events.UpdateProfileCommand{
		NewProfile: *newProfile,
		OldProfile: *oldProfile,
		Type:       events.UpdateProfile,
	}
	return o.commandPublisher.Publish(event)
}

func (o *UpdateProfileOrchestrator) handle(reply *events.UpdateProfileReply) {
	command := events.UpdateProfileCommand{
		NewProfile: reply.NewProfile,
		OldProfile: reply.OldProfile,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *UpdateProfileOrchestrator) nextCommandType(reply events.UpdateProfileReplyType) events.UpdateProfileCommandType {
	switch reply {
	case events.ProfileNotUpdated:
		return events.RollbackUpdatedProfile
	default:
		return events.UnknownCommand
	}
}
