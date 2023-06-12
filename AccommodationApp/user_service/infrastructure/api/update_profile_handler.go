package api

import (
	saga "accommodation_booking/common/saga/messaging"
	events "accommodation_booking/common/saga/update_profile"
	"accommodation_booking/user_service/application"
	"context"
)

type UpdateProfileCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		userService:       userService,
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
	case events.UpdateProfile:
		if command.NewProfile.Username == command.OldProfile.Username {
			return
		}
		_, err := handler.userService.Update(context.TODO(), command.NewProfile.Id, command.NewProfile.Username)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		_, err := handler.userService.Update(context.TODO(), command.NewProfile.Id, command.NewProfile.Username)
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
