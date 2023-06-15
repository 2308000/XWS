package api

import (
	"accommodation_booking/accommodation_service/application"
	saga "accommodation_booking/common/saga/messaging"
	events "accommodation_booking/common/saga/update_profile"
	"context"
)

type UpdateProfileCommandHandler struct {
	accommodationService *application.AccommodationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

func NewUpdateProfileCommandHandler(accommodationService *application.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		accommodationService: accommodationService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
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
		accommodations, err := handler.accommodationService.GetByHost(context.TODO(), command.OldProfile.Id.Hex())
		if err != nil {
			return
		}
		for _, accommodation := range accommodations {
			accommodation.Host.Username = command.NewProfile.Username
			accommodation.Host.PhoneNumber = command.NewProfile.PhoneNumber
			_, err = handler.accommodationService.Update(context.TODO(), accommodation.Id.Hex(), accommodation)
			if err != nil {
				return
			}
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		accommodations, err := handler.accommodationService.GetByHost(context.TODO(), command.NewProfile.Id.Hex())
		if err != nil {
			return
		}
		for _, accommodation := range accommodations {
			accommodation.Host.Username = command.OldProfile.Username
			accommodation.Host.PhoneNumber = command.OldProfile.PhoneNumber
			_, err = handler.accommodationService.Update(context.TODO(), accommodation.Id.Hex(), accommodation)
			if err != nil {
				return
			}
		}
		reply.Type = events.ProfileUpdateRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
