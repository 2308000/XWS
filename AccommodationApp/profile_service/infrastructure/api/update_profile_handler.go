package api

import (
	saga "accommodation_booking/common/saga/messaging"
	events "accommodation_booking/common/saga/update_profile"
	"accommodation_booking/profile_service/application"
	"context"
)

type UpdateProfileCommandHandler struct {
	profileService    *application.ProfileService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(profileService *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		profileService:    profileService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *UpdateProfileCommandHandler) handle(command *events.UpdateProfileCommand) {
	reply := &events.UpdateProfileReply{
		NewProfile: command.NewProfile,
		OldProfile: command.OldProfile,
		Type:       events.UnknownReply,
	}
	switch command.Type {
	case events.RollbackUpdatedProfile:
		oldProfile := command.OldProfile
		err := handler.profileService.RollbackUpdate(context.TODO(), mapAuthProfileToProfile(&oldProfile))
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdateRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
